package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Queue struct {
	next  *Queue
	value int
}

var queueSize int
var tail *Queue

func enqueue(s **Queue, num int) {
	top := new(Queue)
	if top == nil {
		panic("StackOverflow")
	}

	top.value = num
	// Если раньше элементов не было - устанавливаем head и tail на top
	if queueSize == 0 {
		tail = top
		*s = top
	} else {
		// Иначе создаем ссылку предыдущего элемента на текущий и делаем текущий эл. последним
		tail.next = top
		tail = top
	}

	queueSize++
}

func dequeue(s **Queue) int {
	if queueSize == 0 {
		panic("Underflow")
	}
	val := (*s).value
	*s = (*s).next
	queueSize--

	return val
}

func main() {
	sc, out := bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	var elem string
	stck := new(Queue)

	for i := 0; i < n; i++ {
		sc.Scan()
		elem = sc.Text()
		if elem == "-" {
			fmt.Fprintln(out, dequeue(&stck))
		} else {
			var num int
			fmt.Sscanf(elem, "+ %d", &num)
			enqueue(&stck, num)
		}
	}

	out.Flush()
}
