package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) < 1 {
		fmt.Println("Please provide 1 node number: ", len(arguments))
		return
	}

	num, err := strconv.Atoi(arguments[1])
	if err != nil {
		log.Printf("err on strconv Atoi %v", err)
	}

	RaftServer := NewRaftServer(num)

	go RaftServer.ReceiveMssg()

	for i := 0; i < len(RaftServer.nodes); i++ {
		go func(i int) {
			for {
				RaftServer.SendMssg(i)
			}
		}(i)

		//go RaftServer.SendMssg(i)
	}

	go func() {
		for {
			fmt.Print(RaftServer.Receive())
		}
	}()

	for {
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		i := strings.SplitAfterN(input, " ", 2)
		//fmt.Println("i: ", i[0])
		node := i[0]
		node = strings.TrimRight(node, " ")
		nodeNum, err := strconv.Atoi(node)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(nodeNum)
		RaftServer.Send(nodeNum, i[1])
	}

}
