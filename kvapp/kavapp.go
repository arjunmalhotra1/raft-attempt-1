package kvapp

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type KvStore struct {
	kvMap map[string]string
	mu    sync.RWMutex
}

func (kvStore *KvStore) ExecuteCommand(cmd string) string {
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
	case "snapshot":
		filename := fmt.Sprintf("kv-snap-%d.kv", int(time.Now().Unix()))
		kvs.writeSnapshot(filename)
		return filename
	}
	return "error bad command"
}

func (kvs *KvStore) writeSnapshot(filename string) {
	// open output file
	fo, err := os.Create(fmt.Sprintf("%s.txt", filename))
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	enc := gob.NewEncoder(fo)
	err = enc.Encode(kvs.kvMap)
	if err != nil {
		log.Fatal("error encoding kvs map in", err)
	}
}

func NewKVStore() *KvStore {
	kv := new(KvStore)
	kv.kvMap = make(map[string]string)
	return kv
}
