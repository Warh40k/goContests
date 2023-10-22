package main

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestOld(t *testing.T) {
	commands := make([]string, 40e6)
	commands[0] = "create"
	commands[1] = "create"
	var j int
	for j = 2; j < 12; j = j + 3 {
		commands[j] = "insert"
		commands[j+1] = "0"
		commands[j+2] = strconv.Itoa(rand.Intn(10e9))
	}
	for ; j < 22; j = j + 3 {
		commands[j] = "insert"
		commands[j+1] = "1"
		commands[j+2] = strconv.Itoa(rand.Intn(10e9))
	}
	for ; j < 10e6; j = j + 3 {
		commands[j] = "merge"
		commands[j+1] = "0"
		commands[j+2] = "1"
	}
	ExecuteCommandsOlder(commands, j)
}
