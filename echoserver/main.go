package main

import (
	"fmt"
	"net"

	"github.com/arjunmalhotra1/raft/kvapp"
	"github.com/arjunmalhotra1/raft/messaging"
)

func main() {
	// Create a new listener on localhost:8080
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Listening on localhost:8080...")
	kvs := kvapp.NewKVStore()

	// Accept incoming connections indefinitely
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		// Handle each connection in a new goroutine
		go handleConnection(conn, kvs)
	}
}

func handleConnection(conn net.Conn, kvs *kvapp.KvStore) {
	// Read until the connection is closed or an error occurs
	for {
		message, err := messaging.ReceiveMessage(conn)
		if err != nil {
			break
		}
		output := kvs.ExecuteCommand(string(message))
		// Echo the message back to the client
		messaging.SendMessage(conn, []byte(output))

	}
	conn.Close()
}
