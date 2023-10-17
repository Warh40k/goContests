package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Queue[T any] struct {
	size int
	head *QNode[T]
	tail *QNode[T]
}

type QNode[T any] struct {
	value T
	next  *QNode[T]
}

func (q *Queue[T]) push(el T) {
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

func (q *Queue[T]) pop() T {
	if q.size == 0 {
		panic("Underflow")
	}
	val := q.head.value
	q.head = q.head.next
	q.size--

	return val
}

type PriorityMinHeap struct {
	k, heapSize int
	a           [10e6]int
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

func (bh *PriorityMinHeap) insert(i int) {
	bh.heapSize++
	bh.a[bh.heapSize-1] = i
	bh.siftUp(bh.heapSize - 1)
}

func (bh *PriorityMinHeap) getMin() string {
	if bh.heapSize == 0 {
		return "*"
	}
	hmax := bh.a[0]
	bh.a[0] = bh.a[bh.heapSize-1]
	bh.heapSize--
	bh.siftDown(0)

	return strconv.Itoa(hmax)
}

func (bh *PriorityMinHeap) decreaseKey(x, y int) {
	i := 0
	for bh.a[i] != x {
		i++
	}
	bh.a[i] = y
	bh.siftUp(i)
}

func (priors *Queue[T]) findPriority(k int) T {
	var heap = priors.head
	for i := 0; i < k; i++ {
		heap = heap.next
	}
	return heap.value
}

func executeCommands(commands *Queue[string]) *Queue[string] {
	priors := new(Queue[*PriorityMinHeap])
	result := new(Queue[string])

	for commands.size != 0 {

		switch commands.pop() {
		case "create":
			priors.push(new(PriorityMinHeap))
		case "insert":
			k, _ := strconv.Atoi(commands.pop())
			x, _ := strconv.Atoi(commands.pop())
			priors.findPriority(k).insert(x)
		case "extract-min":
			k, _ := strconv.Atoi(commands.pop())
			result.push(priors.findPriority(k).getMin())
		case "decrease-key":
			k, _ := strconv.Atoi(commands.pop())
			x, _ := strconv.Atoi(commands.pop())
			y, _ := strconv.Atoi(commands.pop())
			priors.findPriority(k).decreaseKey(x, y)
		case "merge":
			k, _ := strconv.Atoi(commands.pop())
			m, _ := strconv.Atoi(commands.pop())
			kq, mq := priors.findPriority(k), priors.findPriority(m)
			merged := new(PriorityMinHeap)
			for i := 0; i < kq.heapSize; i++ {
				merged.insert(kq.a[i])
			}
			for i := 0; i < mq.heapSize; i++ {
				merged.insert(mq.a[i])
			}
			priors.push(merged)
		}
	}
	return result
}

func main() {
	out := bufio.NewWriter(os.Stdout)
	commands := new(Queue[string])

	for {
		var cmd string
		_, err := fmt.Scan(&cmd)
		if err == io.EOF {
			break
		}
		commands.push(cmd)
	}

	result := executeCommands(commands).head
	for result != nil {
		if result.next != nil {
			fmt.Fprintln(out, result.value)
		} else {
			fmt.Fprint(out, result.value)
		}
		result = result.next
	}
	out.Flush()
}
