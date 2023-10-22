package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Queue[T any] struct {
	Size     int64
	Head     *QNode[T]
	tail     *QNode[T]
	Iterator *QNode[T]
}

type QNode[T any] struct {
	Value T
	Next  *QNode[T]
}

func (q *Queue[T]) ResetIterator() {
	q.Iterator = q.Head
}

func (q *Queue[T]) Next() T {
	res := q.Iterator
	q.Iterator = q.Iterator.Next
	return res.Value
}

func (q *Queue[T]) Push(el T) {
	top := new(QNode[T])
	if top == nil {
		panic("StackOverflow")
	}

	top.Value = el
	// Если раньше элементов не было - устанавливаем head и tail на top
	if q.Size == 0 {
		q.Head = top
		q.tail = top
	} else {
		// Иначе создаем ссылку предыдущего элемента на текущий и делаем текущий эл. последним
		q.tail.Next = top
		q.tail = top
	}

	q.Size++
}

func (q *Queue[T]) Pop() T {
	if q.Size == 0 {
		panic("Underflow")
	}
	val := q.Head.Value
	q.Head = q.Head.Next
	q.Size--

	return val
}

type PriorityMinHeap struct {
	HeapSize int
	a        [100]int
	Heaped   bool
}

func (bh *PriorityMinHeap) siftUp(i int) {
	for bh.a[i] < bh.a[(i-1)/2] {
		bh.a[i], bh.a[(i-1)/2] = bh.a[(i-1)/2], bh.a[i]
		i = (i - 1) / 2
	}
}

func (bh *PriorityMinHeap) siftDown(i int) {
	for 2*i+1 < bh.HeapSize {
		left, right := 2*i+1, 2*i+2
		j := left
		if right < bh.HeapSize && bh.a[right] < bh.a[left] {
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
	bh.HeapSize++
	bh.a[bh.HeapSize-1] = i
	if bh.Heaped {
		bh.siftUp(bh.HeapSize - 1)
	}
}

func (bh *PriorityMinHeap) getMin() string {
	if bh.HeapSize == 0 {
		return "*"
	}
	if !bh.Heaped {
		bh.build(bh.HeapSize)
		bh.Heaped = true
	}
	hmax := bh.a[0]
	bh.a[0] = bh.a[bh.HeapSize-1]
	bh.HeapSize--
	bh.siftDown(0)

	return strconv.Itoa(hmax)
}

func (bh *PriorityMinHeap) decreaseKey(x, y int) {
	for i := 0; i < bh.HeapSize; i++ {
		if bh.a[i] == x {
			bh.a[i] = y
			bh.siftUp(i)
			break
		}
	}
}

func (bh *PriorityMinHeap) build(N int) {
	bh.HeapSize = N
	for i := bh.HeapSize / 2; i >= 0; i-- {
		bh.siftDown(i)
	}
}

func ExecuteCommandsOlder(commands []string) *Queue[string] {
	priors := make([]*PriorityMinHeap, 10e6)
	result := new(Queue[string])
	lastprior := 0
	for i := 0; commands[i] != ""; i++ {

		switch commands[i] {
		case "create":
			lastprior++
		case "insert":
			k, _ := strconv.Atoi(commands[i+1])
			if priors[k] == nil {
				priors[k] = new(PriorityMinHeap)
			}
			x, _ := strconv.Atoi(commands[i+2])
			priors[k].insert(x)
			i += 2
		case "extract-min":
			k, _ := strconv.Atoi(commands[i+1])
			if priors[k] == nil {
				result.Push("*")
			} else {
				result.Push(priors[k].getMin())
			}
			i += 1
		case "decrease-key":
			k, _ := strconv.Atoi(commands[i+1])
			if priors[k] != nil {
				x, _ := strconv.Atoi(commands[i+2])
				y, _ := strconv.Atoi(commands[i+3])
				priors[k].decreaseKey(x, y)
			}
			i += 3
		case "merge":
			k, _ := strconv.Atoi(commands[i+1])
			m, _ := strconv.Atoi(commands[i+2])
			kq, mq := priors[k], priors[m]
			if priors[k] != nil && priors[m] != nil {
				merged := new(PriorityMinHeap)
				j := 0
				for d := 0; d < kq.HeapSize; d++ {
					merged.a[j] = kq.a[d]
					//merged.insert(kq.a[i])
					j++
				}
				for d := 0; d < mq.HeapSize; d++ {
					merged.a[j] = mq.a[d]
					//merged.insert(mq.a[i])
					j++
				}
				//merged.build(arrmerged, j)
				merged.HeapSize = j
				priors[lastprior] = merged
				lastprior++
				i += 2
			}
		}
	}
	return result
}

func main() {
	out := bufio.NewWriter(os.Stdout)
	commands := make([]string, 40e6)
	var i int
	for i = 0; i < 10e6; i++ {
		_, err := fmt.Scan(&commands[i])
		if err == io.EOF {
			break
		}
	}

	result := ExecuteCommandsOlder(commands).Head
	for result != nil {
		if result.Next != nil {
			fmt.Fprintln(out, result.Value)
		} else {
			fmt.Fprint(out, result.Value)
		}
		result = result.Next
	}
	out.Flush()
}
