package main

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestOld(t *testing.T) {
	commands := new(Queue[string])
	commands.Push("create")
	commands.Push("create")
	for j := 0; j < 10; j++ {
		commands.Push("insert 0" + strconv.Itoa(rand.Intn(10e9)))
		commands.Push("insert 1" + strconv.Itoa(rand.Intn(10e9)))
	}
	for i := 0; i < 10e6; i++ {
		commands.Push("merge 0 1")
	}
	ExecuteCommandsOlder(commands)
	t.Log("End")
}
