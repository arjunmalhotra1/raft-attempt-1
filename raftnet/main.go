package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) < 1 {
		fmt.Println("Please provide 1 node number: ", len(arguments))
		return
	}

	//fmt.Printf("%#v", arguments)

	num, err := strconv.Atoi(arguments[1])
	if err != nil {
		log.Printf("err on strconv Atoi %v", err)
	}

	RaftServer := NewRaftServer(num)

	go RaftServer.ReceiveMssg()

	for i := 0; i < len(RaftServer.nodes); i++ {
		go RaftServer.SendMssg(i)
	}

	go func() {
		fmt.Println(RaftServer.Receive())
	}()

	for {
		bufio.NewReader(os.Stdin)
	}

}
