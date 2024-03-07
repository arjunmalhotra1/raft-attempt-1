package raftlog

type LogEntry struct {
	Term    int
	Command string
}
type RaftLog struct {
	Log []LogEntry
}

func (rl *RaftLog) AppendEntries(prevIndex int, prevTerm int, entries []LogEntry) bool {
	/*
		1. No gaps/holes in the log
		2. assert log[prev_index].term == prev.term
		3. If there are existing entries in the log that have the same position/index as the new entries added,
		but with different terms, they need to be deleted
	*/

	if prevIndex >= len(rl.Log) {
		return false
	}

	if rl.Log[prevIndex].Term != prevTerm {
		return false
	}

	rl.Log = append(rl.Log, entries...)

	return true
}

func (rl *RaftLog) AppendNewCommand(term int, command string) {
	prevIndex := len(rl.Log) - 1
	prevTerm := rl.Log[prevIndex].Term
	rl.AppendEntries(prevIndex, prevTerm, []LogEntry{{term, command}})
}

func NewRaftLog(initialEntries []LogEntry) *RaftLog {
	logs := make([]LogEntry, 1)
	if len(initialEntries) > 0 {
		logs = append(logs, initialEntries...)
	}

	rl := RaftLog{
		Log: logs,
	}
	return &rl
}
