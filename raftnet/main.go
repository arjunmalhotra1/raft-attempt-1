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

	Raftserver := NewRaftServer(num)
	go Raftserver.ReceiveMssg()

	for {
		bufio.NewReader(os.Stdin)
	}

}
