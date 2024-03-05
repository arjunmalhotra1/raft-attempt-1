package raftnet

import (
	"fmt"
	"net"

	"github.com/arjunmalhotra1/raft/messaging"
)

type Node struct {
	nodeNum int
}

func (n *Node) Send(destination int, msg string) {
	// Send a message to a given destination
	// Does not wait for a message to be delivered
	// Does not guarantee that a message will arrive
	// Does not return any kind of error
}

func (n *Node) Receive(destination int, msg string) {
	// Send a message to a given destination
	// Does not specify that has been delivered to us
	// Does not guarantee any particular message order
	// Waits until a message actually arrives if none available
}

func (n *Node) ListenForConnection() {
	// Create a new listener on localhost:8080
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Listening on localhost:8080...")

	// Accept incoming connections indefinitely
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		// Handle each connection in a new goroutine
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	// Read until the connection is closed or an error occurs
	for {
		message, err := messaging.ReceiveMessage(conn)
		if err != nil {
			break
		}
		// Echo the message back to the client
		messaging.SendMessage(conn, message)
	}
	conn.Close()
}

func NewNode(n int) Node {
	return Node{
		nodeNum: n,
	}
}
