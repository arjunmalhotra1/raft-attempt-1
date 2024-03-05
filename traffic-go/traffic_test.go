package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrafficSignal(t *testing.T) {
	signal := NewTrafficSignal()
	signal.setTrafficSignal("G", "R", false, false)

	assert.Equal(t, signal.ewButton, "G")
	assert.Equal(t, signal.nsButton, "R")
}
