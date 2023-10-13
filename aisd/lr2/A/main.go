package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Stack struct {
	next  *Stack
	value int
}

var stackSize int

func push(s **Stack, num int) {
	top := new(Stack)
	if top == nil {
		panic("StackOverflow")
	}
	top.value = num
	top.next = *s
	*s = top
	stackSize++
}

func pop(s **Stack) int {
	if stackSize == 0 {
		panic("Underflow")
	}
	val := (*s).value
	*s = (*s).next
	stackSize--

	return val
}

func main() {
	sc, out := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	var elem string
	stck := new(Stack)

	for i := 0; i < n; i++ {
		sc.Scan()
		elem = sc.Text()
		if elem == "-" {
			fmt.Fprintln(out, pop(&stck))
		} else {
			var num int
			fmt.Sscanf(elem, "+ %d", &num)
			push(&stck, num)
		}
	}

	out.Flush()
}
