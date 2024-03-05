package main

import (
	"fmt"
)

type TrafficSignal struct {
	ewColor  string
	nsColor  string
	ewButton bool
	nsButton bool
	clock    int
}

func (ts *TrafficSignal) String() {
	fmt.Printf("%+v \n", ts)
}

func (ts *TrafficSignal) setTrafficSignal(ewColor, nsColor string, ewButton, nsButton bool) {
	ts.ewColor = ewColor
	ts.nsColor = nsColor
	ts.ewButton = ewButton
	ts.nsButton = nsButton
}

func (ts *TrafficSignal) handleClockTick() {
	ts.clock += 1

	if ts.ewColor == "G" && ts.nsColor == "R" && ts.nsButton == false && ts.clock == 30 {
		ts.ewColor = "Y"
		ts.clock = 0
	}

	if ts.ewColor == "G" && ts.nsColor == "R" && ts.nsButton == true && ts.clock >= 15 {
		ts.ewColor = "Y"
		ts.clock = 0
	}

	if ts.ewColor == "Y" && ts.nsColor == "R" && ts.clock == 5 {
		ts.ewColor = "R"
		ts.nsColor = "G"
		ts.clock = 0
	}

	if ts.ewColor == "R" && ts.nsColor == "G" && ts.nsButton == false && ts.clock == 60 {
		ts.nsColor = "Y"
		ts.clock = 0
	}

	if ts.ewColor == "R" && ts.nsColor == "G" && ts.nsButton == true && ts.clock >= 15 {
		ts.nsColor = "Y"
		ts.clock = 0
	}

	if ts.ewColor == "R" && ts.nsColor == "Y" && ts.clock == 5 {
		ts.ewColor = "G"
		ts.nsColor = "R"
		ts.clock = 0
	}
}

func (ts *TrafficSignal) handleEWButton() {
	ts.ewButton = true
}

func (ts *TrafficSignal) handleNSButton() {
	ts.nsButton = true
}

func NewTrafficSignal() *TrafficSignal {

	ts := TrafficSignal{
		ewColor:  "R",
		nsColor:  "G",
		ewButton: false,
		nsButton: false,
		clock:    0,
	}

	return &ts
}
