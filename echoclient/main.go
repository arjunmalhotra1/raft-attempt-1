package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/arjunmalhotra1/raft/messaging"
)

func main() {
	// Connect to the server on localhost:8080
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Use a bufio.Scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)

	// Read lines from stdin and send them to the server
	for scanner.Scan() {
		message := scanner.Text()
		messaging.SendMessage(conn, []byte(message))
		response, err := messaging.ReceiveMessage(conn)
		if err != nil {
			break
		}
		fmt.Println(string(response))
	}
}
