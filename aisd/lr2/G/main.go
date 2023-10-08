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
	sum        int
	head, tail *Node
}

func (deque *Deque) pushHead(num int) {
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
	(*deque).sum += num
	(*deque).size++
}

func (deque *Deque) pushTail(num int) {
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
	(*deque).sum += num
	(*deque).size++
}

func (deque *Deque) popHead() int {
	if (*deque).size == 0 {
		return 0
	}
	val := (*deque).head.value
	(*deque).head = (*deque).head.next
	if deque.head != nil {
		deque.head.prev = nil
	}
	(*deque).sum -= val
	(*deque).size--

	return val
}

func (deque *Deque) popTail() int {
	if (*deque).size == 0 {
		return 0
	}
	val := (*deque).tail.value
	deque.tail = (*deque).tail.prev
	if deque.tail != nil {
		deque.tail.next = nil
	}

	(*deque).sum -= val
	(*deque).size--

	return val
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var N int
	fmt.Fscan(in, N)
	var arr = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &arr[i])
	}

	fmt.Fprintln(out, getSquare(N, arr))
	out.Flush()
}

func getSquare(N int, arr []int) int {
	sides := make([]Deque, 4)
	var i, result int

	//for i := 0; i < N; i++ {
	//	sides[0].pushTail(arr[i])
	//}
	//sum = sides[0].sum
	//cursum = sum

	for {
		val := sides[i].popTail()
		sides[(i+1)%4].pushHead(val)
		if sides[i].sum > cursum/2 {
			if i == 3 {
				break
			}
		} else {
			if i == 3 {
				cursum = sum
			} else {
				cursum -= sides[i].sum
			}
			i = (i + 1) % 4
		}
	}

	return result
	//for i := 0; i < 4; i++ {
	//	result += sides[i].sum
	//}
}
