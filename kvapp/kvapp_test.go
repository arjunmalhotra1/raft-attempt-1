package kvapp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteCommand(t *testing.T) {
	kvStore := NewKVStore()
	res := kvStore.ExecuteCommand("get hello")
	assert.Equal(t, res, "")
	res = kvStore.ExecuteCommand("put hello world")
	assert.Equal(t, res, "ok")
	res = kvStore.ExecuteCommand("get hello")
	assert.Equal(t, res, "world")
}
