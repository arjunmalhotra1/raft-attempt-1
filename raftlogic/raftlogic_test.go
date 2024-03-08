package raftlogic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBecomeLeader(t *testing.T) {
	server0 := NewRaftServerLogic(0, 2)
	server1 := NewRaftServerLogic(1, 2)
	server0.BecomeLeader()
	Message := Message{
		source:      0,
		destination: 0,
		term:        1,
	}
	upf := UpdateFollowers{
		Message: Message,
	}
	req := ApplicationRequest{
		Message: Message,
		command: "set x 42",
	}
	server0.HandleMessage(req)
	//fmt.Println("len log.Log: ", len(server0.log.Log))
	server0.HandleMessage(upf)

	server1.HandleMessage(server0.outgoingAppendEntries[0])
	server0.HandleMessage(server1.outgoingAppendEntriesResponse[0])

	assert.Equal(t, server0.log.Log, server1.log.Log)
	assert.Equal(t, server0.nextIndex[1], 2)

}
