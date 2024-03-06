package main

import "github.com/arjunmalhotra1/raft/raftlog"

type RaftServerLogic struct {
	nodeNum     int
	clusterSize int
	log         *raftlog.RaftLog
	currentTerm int
	commitIndex int
}

type Message struct {
	source      int
	destination int
	term        int
}

type AppenEntries struct {
	//term      int
	msg       Message
	prevIndex int
	prevTerm  int
	entries   []raftlog.LogEntry
}

type AppenEntriesResponse struct {
	success bool
}

func (rsl *RaftServerLogic) updateFollowers() {
	for i := 0; i < rsl.clusterSize; i++ {
		if i != rsl.nodeNum {
			msg := AppenEntries{}
			send(msg)
		}
	}
}

func (rsl *RaftServerLogic) HandleAppendEntries(msg AppenEntries) {
	success := rsl.log.AppendEntries(msg.prevIndex, msg.prevTerm, msg.entries)
	response := AppenEntriesResponse{}
	send(response)
}

func (rsl *RaftServerLogic) HandleAppendEntriesResponse(msg AppenEntriesResponse) {
	if msg.success {
		// It worked!
	} else {
		// It Failed!
		// Now what?!?!
	}
}

func NewRaftServerLogic(nodeNum, clusterSize int) *RaftServerLogic {
	rsl := RaftServerLogic{
		nodeNum:     nodeNum,
		clusterSize: clusterSize,
		log:         raftlog.NewRaftLog([]raftlog.LogEntry{}),
		currentTerm: 1,
		commitIndex: 0,
	}
	return &rsl
}
