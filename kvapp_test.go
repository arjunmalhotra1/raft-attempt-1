package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteCommand(t *testing.T) {
	kvStore := NewKVStore()
	res := kvStore.executeCommand("get hello")
	assert.Equal(t, res, "")
	res = kvStore.executeCommand("put hello world")
	assert.Equal(t, res, "ok")
	res = kvStore.executeCommand("get hello")
	assert.Equal(t, res, "world")
}
