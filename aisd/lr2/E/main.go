package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	size int
	top  *SNode
}

type SNode struct {
	next  *SNode
	value int
}

func (s *Stack) push(num int) {
	temp := new(SNode)

	temp.value = num
	temp.next = s.top
	s.top = temp

	s.size++
}

func (s *Stack) pop() int {
	val := s.top.value
	s.top = s.top.next

	s.size--

	return val
}

func (s *Stack) isEmpty() bool {
	if s.size == 0 {
		return true
	}
	return false
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscan(in, &n)
	var input = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &input[i])
	}

	fmt.Fprintln(out, parseCalendar(n, input))
	out.Flush()
}

func parseCalendar(n int, input []int) string {
	result := make([]string, n)

	sSearch := new(Stack)

	for i := 0; i < n; i++ {
		result[i] = strconv.Itoa(-1)
		for !sSearch.isEmpty() {
			k := sSearch.top.value

			if input[i] > input[k] {
				result[k] = strconv.Itoa(i - k)
				sSearch.pop()
			} else {
				break
			}
		}
		sSearch.push(i)
	}

	return strings.Join(result, " ")
}
