package main

import (
	"net"
	"time"
)

var eventChan chan string
var ts *TrafficSignal

func main() {
	eventChan = make(chan string)
	ts = NewTrafficSignal()

	go generateClockTicks()
	go runEvents()
	go watchButtons()

}

func watchButtons() {

}

func generateClockTicks() {
	for {
		time.Sleep(1 * time.Second)
		eventChan <- "tick"
	}

}

func runEvents() {
	sock, err := net.Listen("udp", "localhost:8080")
	for {
		event := <-eventChan
		switch event {
		case "tick":
			ts.handleClockTick()
		case "ns_button":
			ts.handleNSButton()
		case "ew_button":
			ts.handleEWButton()
		}

		sock
	}
}
