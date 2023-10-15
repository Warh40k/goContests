package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Stack[T any] struct {
	size int
	top  *SNode[T]
}

type SNode[T any] struct {
	next  *SNode[T]
	value T
}

func (s *Stack[T]) push(el T) {
	temp := new(SNode[T])

	temp.value = el
	temp.next = s.top
	s.top = temp

	s.size++
}

func (s *Stack[T]) pop() T {
	val := s.top.value
	s.top = s.top.next

	s.size--

	return val
}

func (s *Stack[T]) isEmpty() bool {
	if s.size == 0 {
		return true
	}
	return false
}

type Queue[T any] struct {
	size int
	head *QNode[T]
	tail *QNode[T]
}

type QNode[T any] struct {
	value T
	next  *QNode[T]
}

func (q *Queue[T]) enqueue(el T) {
	top := new(QNode[T])
	if top == nil {
		panic("StackOverflow")
	}

	top.value = el
	// Если раньше элементов не было - устанавливаем head и tail на top
	if q.size == 0 {
		q.head = top
		q.tail = top
	} else {
		// Иначе создаем ссылку предыдущего элемента на текущий и делаем текущий эл. последним
		q.tail.next = top
		q.tail = top
	}

	q.size++
}

func (q *Queue[T]) dequeue() T {
	if q.size == 0 {
		panic("Underflow")
	}
	val := q.head.value
	q.head = q.head.next
	q.size--

	return val
}

type PriorityMinHeap struct {
	a        []int
	heapSize int
	k        int
}

func (bh *PriorityMinHeap) siftUp(i int) {
	for bh.a[i] < bh.a[(i-1)/2] {
		bh.a[i], bh.a[(i-1)/2] = bh.a[(i-1)/2], bh.a[i]
		i = (i - 1) / 2
	}
}

func (bh *PriorityMinHeap) siftDown(i int) {
	for 2*i+1 < bh.heapSize {
		left, right := 2*i+1, 2*i+2
		j := left
		if right < bh.heapSize && bh.a[right] < bh.a[left] {
			j = right
		}
		if bh.a[i] <= bh.a[j] {
			break
		}
		bh.a[i], bh.a[j] = bh.a[j], bh.a[i]
		i = j
	}
}

func (bh *PriorityMinHeap) getMin(k int) string {
	if bh.heapSize == 0 {
		return "*"
	}
	hmax := bh.a[0]
	bh.a[0] = bh.a[bh.heapSize-1]
	bh.heapSize--
	bh.siftDown(0)

	return strconv.Itoa(hmax)
}

func insertToKQueue(k, x int) {

}

func executeCommands(commands Queue[string]) Queue[int] {
	priorities := new(Stack[*PriorityMinHeap])

	for commands.size != 0 {
		command := strings.Split(commands.dequeue(), " ")
		switch command[0] {
		case "create":
			priorities.push(new(PriorityMinHeap))
		case "insert":

		}
	}
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	commands := new(Queue[string])

	for {
		cmd, err := in.ReadString('\n')
		if err == io.EOF {
			break
		}
		commands.enqueue(cmd)
	}

	fmt.Fprintln(out, executeCommands(commands))
	out.Flush()
}
