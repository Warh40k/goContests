package main

import (
	"math/rand"
	"strconv"
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
		commands1 := new(Queue[string])
		commands1.push("create")
		commands1.push("create")
		commands2 := new(Queue[string])
		commands2.push("create")
		commands2.push("create")
		cmdNames := []string{"create", "merge", "insert", "extract-min"}
		var i int
		for i = 2; i < 7; i++ {
			command := cmdNames[rand.Intn(2)]
			if command == "merge" {
				a, b := strconv.Itoa(rand.Intn(i)), strconv.Itoa(rand.Intn(i))
				//command = strings.Join([]string{command, a, b}, " ")
				commands1.push(command)
				commands1.push(a)
				commands1.push(b)
				commands2.push(a)
				commands2.push(b)
				commands2.push(command)
			} else {
				commands1.push(command)
				commands2.push(command)
			}
		}
		for j := 0; j < 30; j++ {
			command := cmdNames[rand.Intn(4)]
			if command == "merge" {
				a, b := strconv.Itoa(rand.Intn(i)), strconv.Itoa(rand.Intn(i))
				//command += strings.Join([]string{command, a, b}, " ")
				commands1.push(command)
				commands1.push(a)
				commands1.push(b)
				commands2.push(command)
				commands2.push(a)
				commands2.push(b)
			} else if command == "insert" {
				//command = strings.Join([]string{command, strconv.Itoa(rand.Intn(i + 1)), strconv.Itoa(30)}, " ")
				commands1.push(command)
				commands1.push(strconv.Itoa(rand.Intn(i)))
				commands1.push(strconv.Itoa(rand.Intn(30)))
				commands2.push(command)
				commands2.push(strconv.Itoa(rand.Intn(i)))
				commands2.push(strconv.Itoa(rand.Intn(30)))

			} else if command == "extract-min" {
				//command = command + " " + strconv.Itoa(rand.Intn(i+1))
				commands1.push(command)
				commands1.push(strconv.Itoa(rand.Intn(i)))
				commands2.push(command)
				commands2.push(strconv.Itoa(rand.Intn(i)))
			} else if command == "decrease-key" {
				old := rand.Intn(rand.Intn(30)) + 1
				neww := rand.Intn(old + 1)
				//command = command + strconv.Itoa(rand.Intn(i+1)) + " " + strconv.Itoa(old) + " " + strconv.Itoa(neww)
				commands1.push(command)
				commands1.push(strconv.Itoa(rand.Intn(i)))
				commands1.push(strconv.Itoa(old))
				commands1.push(strconv.Itoa(neww))
				commands2.push(command)
				commands2.push(strconv.Itoa(rand.Intn(i)))
				commands2.push(strconv.Itoa(old))
				commands2.push(strconv.Itoa(neww))
			}

		}
		comarr := make([]string, commands1.size)
		com := commands1.head
		for i := 0; i < commands1.size; i++ {
			comarr[i] = com.value
			com = com.next
		}

		//coms := strings.Join(comarr, "\n")

		result1 := executeCommands(commands1).head
		//result2 := executeCommandsOld(commands2).head

		for i := 0; i < 100 && result1 != nil; i++ {

		}
	}

}
