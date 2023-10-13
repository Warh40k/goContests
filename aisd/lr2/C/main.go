package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type Deque struct {
	size       int
	head, tail *Node
}

func pushHead(deque *Deque, num int) {
	temp := new(Node)
	if temp == nil {
		panic("StackOverflow")
	}
	temp.value = num

	if (*deque).size == 0 {
		(*deque).head = temp
		(*deque).tail = temp
	} else {
		(*deque).head.prev = temp
		temp.next = (*deque).head
		(*deque).head = temp
	}
	(*deque).size++
}

func pushTail(deque *Deque, num int) {
	node := new(Node)
	if node == nil {
		panic("StackOverflow")
	}
	node.value = num

	if (*deque).size == 0 {
		(*deque).head = node
		(*deque).tail = node
	} else {
		(*deque).tail.next = node
		node.prev = (*deque).tail
		(*deque).tail = node
	}
	(*deque).size++
}

func popHead(deque *Deque) int {
	if (*deque).size == 0 {
		return 0
	}
	val := (*deque).head.value
	(*deque).head = (*deque).head.next
	if deque.head != nil {
		deque.head.prev = nil
	}
	(*deque).size--

	return val
}

func popTail(deque *Deque) int {
	if (*deque).size == 0 {
		return 0
	}
	val := (*deque).tail.value
	deque.tail = (*deque).tail.prev
	if deque.tail != nil {
		deque.tail.next = nil
	}
	(*deque).size--

	return val
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var M, N, lev int
	fmt.Fscan(in, &M, &N)

	loot := new(Deque)

	// Заполнение сокровищницы
	for i := 0; i < M; i++ {
		fmt.Fscan(in, &lev)
		pushTail(loot, lev)
	}

	rucksack := fillRucksack(N, loot).tail

	for rucksack != nil {
		fmt.Fprint(out, rucksack.value, " ")
		rucksack = rucksack.prev
	}
	out.Flush()
}

func fillRucksack(N int, loot *Deque) *Deque {
	rucksack := new(Deque)
	var found [3]int

	for loot.size != 0 {
		if rucksack.size < N {
			pushHead(rucksack, popHead(loot))
		} else {
			found[0], found[1], found[2] = popTail(rucksack), popHead(rucksack), popHead(loot)
			max1, max2 := getTop2(found)
			pushHead(rucksack, max2)
			pushHead(rucksack, max1)
		}
	}

	return rucksack
}

func getTop2(arr [3]int) (int, int) {
	var max1, max2 int
	for i := 0; i < 3; i++ {
		if arr[i] > max1 {
			max2 = max1
			max1 = arr[i]
		} else if arr[i] > max2 {
			max2 = arr[i]
		}
	}
	return max1, max2
}
