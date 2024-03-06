package raftlog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendEntries(t *testing.T) {
	logs := NewRaftLog(nil)

	// Index is bad. Out of range. Would create a gap
	res := logs.AppendEntries(10, 0, []LogEntry{{term: 1, command: "x"}})
	assert.Equal(t, res, false)

	// Index is good, but prev_term doesn't match up right
	logs = NewRaftLog([]LogEntry{{term: 1, command: "x"}, {term: 2, command: "y"}})
	res = logs.AppendEntries(2, 1, []LogEntry{{term: 1, command: "x"}})
	assert.Equal(t, res, false)

	// Try a successful log append
	res = logs.AppendEntries(2, 2, []LogEntry{{term: 2, command: "z"}})
	assert.Equal(t, res, true)

	// Try a repeated log append.  It should work and log should be unchanged by it
	res = logs.AppendEntries(2, 2, []LogEntry{{term: 2, command: "z"}})
	testLogs := []LogEntry{{term: 1, command: "x"}, {term: 2, command: "y"}, {term: 2, command: "z"}}
	assert.Equal(t, res, true)
	assert.ElementsMatch(t, logs.Log[1:], testLogs)
}
