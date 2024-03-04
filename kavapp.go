package main

import (
	"strings"
	"sync"
)

type kvStore struct {
	kvMap map[string]string
	mu    sync.RWMutex
}

func (kvStore *kvStore) executeCommand(cmd string) string {
	cmds := strings.Split(cmd, " ")
	kvs := *kvStore
	switch cmds[0] {
	case "get":
		defer kvs.mu.RUnlock()
		kvs.mu.RLock()
		return kvs.kvMap[cmds[1]]

	case "put":
		defer kvs.mu.Unlock()
		kvs.mu.Lock()
		key := cmds[1]
		value := cmds[2]
		kvs.kvMap[key] = value
		return "ok"

	case "delete":
		defer kvs.mu.Unlock()
		kvs.mu.Lock()
		key := cmds[1]
		delete(kvs.kvMap, key)
		return "deleted"
		// default:
		// 	return "error bad command"
	}
	return "error bad command"
}

func NewKVStore() *kvStore {
	kv := new(kvStore)
	kv.kvMap = make(map[string]string)
	return kv
}
