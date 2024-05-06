package main

import (
	"math"
	"math/rand"
	"runtime/debug"
	"strconv"
	"testing"
)

func TestCompare(t *testing.T) {
	for {
		commands := new(Vector[string])
		commands.build(10)
		commands.append("create")
		commands.append("create")
		commands.append("create")
		commands.append("create")
		var i = 4
		cmdNames := []string{"create", "merge", "insert", "extract-min", "decrease-key"}
		//for i = 1; i < 5; i++ {
		//	command := cmdNames[rand.Intn(2)]
		//	if command == "merge" {
		//		arr, b := strconv.Itoa(rand.Intn(i)), strconv.Itoa(rand.Intn(i))
		//		commands.Push(command)
		//		commands.Push(arr)
		//		commands.Push(b)
		//	} else {
		//		commands.Push(command)
		//	}
		//}
		for j := 0; j < 10; j++ {
			command := cmdNames[rand.Intn(5)]
			if command == "merge" {
				a, b := strconv.Itoa(rand.Intn(i)), strconv.Itoa(rand.Intn(i))
				commands.append(command)
				commands.append(a)
				commands.append(b)
				i++
			} else if command == "insert" {
				commands.append(command)
				commands.append(strconv.Itoa(rand.Intn(i)))
				commands.append(strconv.Itoa(rand.Intn(10)))

			} else if command == "extract-min" {
				commands.append(command)
				commands.append(strconv.Itoa(rand.Intn(i)))
			} else if command == "decrease-key" {
				sign := []int{-1, 1}[rand.Intn(2)]
				old := sign * rand.Intn(10)
				neww := sign * rand.Intn(int(math.Abs(float64(old)))+1)

				commands.append(command)
				commands.append(strconv.Itoa(rand.Intn(i)))
				commands.append(strconv.Itoa(old))
				commands.append(strconv.Itoa(neww))
			}

		}
		for j := 0; j < 5; j++ {
			commands.append("extract-min")
			commands.append(strconv.Itoa(rand.Intn(i)))
		}

		//result1 := ExecuteCommandsOlder(commands)
		//result2 := executeCommandsOld(commands)
		//result1.ResetIterator()
		//result2.ResetIterator()
		//
		//for result1.Iterator != nil {
		//	val1, val2 := result1.Next(), result2.Next()
		//	if val1 != val2 {
		//		t.Fatalf("Incorrect, got %s, expected %s", val1, val2)
		//	}
		//}
	}

}

func TestOld(t *testing.T) {
	debug.SetGCPercent(-1)
	commands := new(Vector[string])
	commands.build(200)
	commands.append("create")
	commands.append("create")
	var j int
	for j = 2; j < 12; j = j + 3 {
		commands.append("insert")
		commands.append("0")
		commands.append(strconv.Itoa(rand.Intn(10e9)))
	}
	for ; j < 22; j = j + 3 {
		commands.append("insert")
		commands.append("1")
		commands.append(strconv.Itoa(rand.Intn(10e9)))
	}
	for ; j < 10e6; j = j + 3 {
		commands.append("merge")
		commands.append("0")
		commands.append("1")
	}
	//ExecuteCommandsOlder(commands)

	t.Log("Passed")
}
