package main

import (
	"goContests/aisd/lr3/B/qu"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestHeap(t *testing.T) {
	heap := new(MinHeap)
	heap.insert(5)
	heap.insert(4)
	heap.insert(3)
	heap.insert(2)
	heap.insert(2)
	heap.decreaseKey(5, 3)

	espected := []int{2, 2, 3, 3, 4}
	for i := 0; i < len(espected); i++ {
		val, _ := strconv.Atoi(heap.getMin())
		if espected[i] != val {
			t.Errorf("INcorrect, expected %d, got %d", espected[i], val)
		}
	}

}

func TestAlgo(t *testing.T) {
	for {
		commands := new(qu.Queue[string])
		commands.Push("create")
		cmdNames := []string{"create", "merge", "insert", "extract-min", "decrease-key"}
		var i int
		for i = 1; i < 5; i++ {
			command := cmdNames[rand.Intn(2)]
			if command == "merge" {
				a, b := strconv.Itoa(rand.Intn(i)), strconv.Itoa(rand.Intn(i))
				commands.Push(command)
				commands.Push(a)
				commands.Push(b)
			} else {
				commands.Push(command)
			}
		}
		for j := 0; j < 20; j++ {
			command := cmdNames[rand.Intn(5)]
			if command == "merge" {
				a, b := strconv.Itoa(rand.Intn(i)), strconv.Itoa(rand.Intn(i))
				commands.Push(command)
				commands.Push(a)
				commands.Push(b)
			} else if command == "insert" {
				commands.Push(command)
				commands.Push(strconv.Itoa(rand.Intn(i)))
				commands.Push(strconv.Itoa(-1 * rand.Intn(2) * rand.Intn(10)))

			} else if command == "extract-min" {
				commands.Push(command)
				commands.Push(strconv.Itoa(rand.Intn(i)))
			} else if command == "decrease-key" {
				sign := []int{-1, 1}[rand.Intn(2)]
				old := sign * rand.Intn(10)
				neww := sign * rand.Intn(int(math.Abs(float64(old)))+1)

				commands.Push(command)
				commands.Push(strconv.Itoa(rand.Intn(i)))
				commands.Push(strconv.Itoa(old))
				commands.Push(strconv.Itoa(neww))
			}

		}
		for j := 0; j < 100; j++ {
			commands.Push("extract-min " + strconv.Itoa(rand.Intn(i)))
		}
		comarr := make([]string, commands.Size)
		com := commands.Head
		for j := int64(0); j < commands.Size; j++ {
			comarr[j] = com.Value
			com = com.Next
		}

		coms := strings.Join(comarr, "\n")
		reflect.TypeOf(coms)

		result1 := executeCommands(commands)
		result2 := executeCommandsOld(commands)
		result1.ResetIterator()
		result2.ResetIterator()

		for result1.Iterator != nil {
			val1, val2 := result1.Next(), result2.Next()
			if val1 != val2 {
				t.Fatalf("Incorrect, got %s, expected %s", val1, val2)
			}
		}
		t.Log("Passed")
	}

}
