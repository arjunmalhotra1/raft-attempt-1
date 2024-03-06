package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/arjunmalhotra1/raft/raftconfig"
)

type RaftServer struct {
	nodeNum  int
	nodes    []raftconfig.Node
	inbox    chan string
	outboxes []chan string
}

func (n *RaftServer) Send(destination int, msg string) {
	// Send a message to a given destination
	// Does not wait for a message to be delivered
	// Does not guarantee that a message will arrive
	// Does not return any kind of error
	n.outboxes[destination] <- msg
}

func (n *RaftServer) Receive() string {
	// Send a message to a given destination
	// Does not specify that has been delivered to us
	// Does not guarantee any particular message order
	// Waits until a message actually arrives if none available
	return <-n.inbox
}

func (n *RaftServer) ReceiveMssg() {
	// Create a new listener on localhost:8080
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", n.nodes[n.nodeNum].Port))
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Printf("Listening on localhost:%d...", n.nodes[n.nodeNum].Port)

	// Accept incoming connections indefinitely
	for {
		fmt.Println("Waiting for connection")
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		fmt.Println("After listner Accept()")

		reader := bufio.NewReader(conn)
		message, err := reader.ReadString('\n')
		if err != nil {
			// Handle error or break loop if connection is closed
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading message: %v\n", err)
			break
		}

		fmt.Println("After message", string(message))

		n.inbox <- string(message)
		fmt.Println("message: ", message)
		// Handle each connection in a new goroutine
		//go handleConnection(conn)
	}

}

func (n *RaftServer) SendMssg(destNodeNum int) {

	// Connect to the server on localhost:8080
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", n.nodes[destNodeNum].Port))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	message := <-n.outboxes[destNodeNum]
	// First, we need to send the length of the message as a 4-byte integer
	// Convert the message length to a 4-byte binary representation
	length := make([]byte, 4)
	binary.BigEndian.PutUint32(length, uint32(len(message)))

	// Write the length to the connection
	if _, err := conn.Write(length); err != nil {
		log.Printf("error on Write length %v \n", err)
	}

	// Then, write the message itself to the connection
	if _, err := conn.Write([]byte(message)); err != nil {
		log.Printf("error on Write message %v \n", err)
	}

	fmt.Printf("message sent to node %d \n", n.nodeNum)

}

func NewRaftServer(n int) RaftServer {
	fmt.Printf("I am node num %d\n", n)
	in := make(chan string)
	listNodes := raftconfig.NewNodes()
	out := make([]chan string, len(listNodes))
	return RaftServer{
		nodeNum:  n,
		inbox:    in,
		outboxes: out,
		nodes:    listNodes,
	}
}
