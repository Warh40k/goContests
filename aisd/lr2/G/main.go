package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	size int
	top  *SNode
}

type SNode struct {
	next  *SNode
	value int64
}

func (s *Stack) push(num int64) {
	temp := new(SNode)

	temp.value = num
	temp.next = s.top
	s.top = temp

	s.size++
}

func (s *Stack) pop() int64 {
	val := s.top.value
	s.top = s.top.next

	s.size--

	return val
}

func (q *Stack) isEmpty() bool {
	if q.top == nil {
		return true
	}
	return false
}

func (q *Stack) peek() int64 {
	if q.top == nil {
		return -1
	}
	return q.top.value
}

func (q *Stack) clear() {
	for !q.isEmpty() {
		q.pop()
	}
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var n int64
	fmt.Fscan(in, &n)
	var input = make([]int64, n)

	for i := int64(0); i < n; i++ {
		fmt.Fscan(in, &input[i])
	}

	fmt.Fprintln(out, getMaxSquare(n, input))
	out.Flush()
}

func getMaxSquare(N int64, boards []int64) int64 {
	var square int64
	lSide := make([]int64, N)
	rSide := make([]int64, N)
	lengths := new(Stack)

	for i := int64(0); i < N; i++ {
		for !lengths.isEmpty() && boards[i] <= boards[lengths.peek()] {
			lengths.pop()
		}
		if !lengths.isEmpty() {
			lSide[i] = lengths.peek()
		} else {
			lSide[i] = -1
		}
		lengths.push(i)
	}
	lengths.clear()
	for i := N - 1; i >= 0; i-- {
		for !lengths.isEmpty() && boards[i] <= boards[lengths.peek()] {
			lengths.pop()
		}
		if !lengths.isEmpty() {
			rSide[i] = lengths.peek()
		} else {
			rSide[i] = N
		}
		lengths.push(i)
	}

	for i := int64(0); i < N; i++ {
		tempSquare := (rSide[i] - lSide[i] - 1) * boards[i]
		if tempSquare > square {
			square = tempSquare
		}
	}
	return square
}
