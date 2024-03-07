package raftlog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendEntries(t *testing.T) {
	logs := NewRaftLog(nil)

	// Index is bad. Out of range. Would create a gap
	res := logs.AppendEntries(10, 0, []LogEntry{{Term: 1, Command: "x"}})
	assert.Equal(t, res, false)

	// Index is good, but prev_term doesn't match up right
	logs = NewRaftLog([]LogEntry{{Term: 1, Command: "x"}, {Term: 2, Command: "y"}})
	res = logs.AppendEntries(2, 1, []LogEntry{{Term: 1, Command: "x"}})
	assert.Equal(t, res, false)

	// Try a successful log append
	res = logs.AppendEntries(2, 2, []LogEntry{{Term: 2, Command: "z"}})
	assert.Equal(t, res, true)

	// Try a repeated log append.  It should work and log should be unchanged by it
	res = logs.AppendEntries(2, 2, []LogEntry{{Term: 2, Command: "z"}})
	testLogs := []LogEntry{{Term: 1, Command: "x"}, {Term: 2, Command: "y"}, {Term: 2, Command: "z"}}
	assert.Equal(t, res, true)
	assert.ElementsMatch(t, logs.Log[1:], testLogs)
}
