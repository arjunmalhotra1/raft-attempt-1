package main

import (
	"fmt"
	"sort"

	"github.com/arjunmalhotra1/raft/raftlog"
)

type Message struct {
	source      int
	destination int
	term        int
}

type AppendEntries struct {
	Message   Message
	prevIndex int
	prevTerm  int
	entries   []raftlog.LogEntry
}

type AppendEntriesResponse struct {
	Message    Message
	success    bool
	matchIndex int
}

type ApplicationRequest struct {
	Message Message
	command string
}

type UpdateFollowers struct {
	Message Message
}

type RaftServerLogic struct {
	nodeNum                       int
	clusterSize                   int
	role                          string
	log                           *raftlog.RaftLog
	currentTerm                   int
	commitIndex                   int
	nextIndex                     []int
	matchIndex                    []int
	outgoingAppendEntriesResponse []AppendEntriesResponse
	outgoingAppendEntries         []AppendEntries
}

func (rsl *RaftServerLogic) sendAppendEntriesResponse(ae AppendEntriesResponse) {
	rsl.outgoingAppendEntriesResponse = append(rsl.outgoingAppendEntriesResponse, ae)
}

func (rsl *RaftServerLogic) sendAppendEntries(ae AppendEntries) {
	rsl.outgoingAppendEntries = append(rsl.outgoingAppendEntries, ae)
}

func (rsl *RaftServerLogic) becomeLeader() {
	for i := 0; i < rsl.clusterSize; i++ {
		rsl.nextIndex[i] = len(rsl.log.Log)
	}

	for i := 0; i < rsl.clusterSize; i++ {
		rsl.matchIndex[i] = 0
	}

	rsl.role = "LEADER"
}

func (rsl *RaftServerLogic) HandleMessage(msg any) []AppendEntriesResponse {
	rsl.outgoingAppendEntriesResponse = make([]AppendEntriesResponse, 0)
	switch v := msg.(type) {
	case AppendEntries:
		rsl.HandleAppendEntries(v)
	case AppendEntriesResponse:
		rsl.HandleAppendEntriesResponse(v)
	case ApplicationRequest:
	case UpdateFollowers:
	}
	return rsl.outgoingAppendEntriesResponse
}

func (rsl *RaftServerLogic) HandleAppendEntries(msg AppendEntries) {
	success := rsl.log.AppendEntries(msg.prevIndex, msg.prevTerm, msg.entries)
	response := AppendEntriesResponse{
		Message: Message{source: rsl.nodeNum,
			destination: msg.Message.destination,
			term:        rsl.currentTerm},
		success:    success,
		matchIndex: msg.prevIndex + len(msg.entries),
	}
	rsl.sendAppendEntriesResponse(response)
}

func (rsl *RaftServerLogic) HandleAppendEntriesResponse(msg AppendEntriesResponse) {
	if msg.success {
		// It worked!
		rsl.nextIndex[msg.Message.source] = msg.matchIndex + 1
		rsl.matchIndex[msg.Message.source] = msg.matchIndex + 1
		sort.Ints(rsl.matchIndex)
		newCommitIndex := rsl.clusterSize / 2
		if newCommitIndex > rsl.commitIndex {
			fmt.Println("COMMITTING:", rsl.log.Log[rsl.commitIndex+1:newCommitIndex])
		}
	} else {
		// It Failed!
		// Now what?!?!
		rsl.nextIndex[msg.Message.source] -= 1
	}
}

func (rsl *RaftServerLogic) HandleApplicationRequest(command string) {
	if rsl.role != "LEADER" {
		panic("I am not a leader but the application submitted a command")
	}

	rsl.log.AppendNewCommand(rsl.currentTerm, command)
	rsl.matchIndex[rsl.nodeNum] = len(rsl.log.Log) - 1
}

func (rsl *RaftServerLogic) HandleUpdateFollowers() {
	if rsl.role != "LEADER" {
		panic("I am not the leader, I cannot update the followers")
	}

	for follower := 0; follower < rsl.clusterSize; follower++ {
		if follower != rsl.nodeNum {
			prevIndex := rsl.nextIndex[follower] - 1
			prevTerm := rsl.log.Log[prevIndex].Term
			// This could give index out of bounds in case of a heartbeat/blank message
			entries := rsl.log.Log[prevIndex+1:]
			msg := Message{source: rsl.nodeNum,
				destination: follower,
				term:        rsl.currentTerm}
			ae := AppendEntries{Message: msg,
				prevIndex: prevIndex,
				prevTerm:  prevTerm,
				entries:   entries}
			rsl.sendAppendEntries(ae)
		}
	}
}

func NewRaftServerLogic(nodeNum, clusterSize int) *RaftServerLogic {
	rsl := RaftServerLogic{
		nodeNum:                       nodeNum,
		clusterSize:                   clusterSize,
		log:                           raftlog.NewRaftLog([]raftlog.LogEntry{}),
		currentTerm:                   1,
		nextIndex:                     make([]int, clusterSize),
		role:                          "FOLLOWER",
		outgoingAppendEntriesResponse: []AppendEntriesResponse{},
	}
	return &rsl
}
