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
	k, HeapSize int
	a           [100]int
	Heaped      bool
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
	bh.Heaped = false
	bh.HeapSize++
	bh.a[bh.HeapSize-1] = i
	//bh.siftUp(bh.HeapSize - 1)
}

func (bh *PriorityMinHeap) getMin() string {
	if bh.HeapSize == 0 {
		return "*"
	}
	if !bh.Heaped {
		bh.build(bh.a, bh.HeapSize)
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
			bh.Heaped = false
			//bh.siftUp(i)
			break
		}
	}
}

func (bh *PriorityMinHeap) build(arr [100]int, N int) {
	bh.a = arr
	bh.HeapSize = N
	for i := bh.HeapSize / 2; i >= 0; i-- {
		bh.siftDown(i)
	}
}

func ExecuteCommandsOlder(commands *Queue[string]) *Queue[string] {
	priors := make([]*PriorityMinHeap, 10e6)
	result := new(Queue[string])
	commands.ResetIterator()
	lastprior := 0
	for commands.Iterator != nil {

		switch commands.Next() {
		case "create":
			priors[lastprior] = new(PriorityMinHeap)
			lastprior++
		case "insert":
			k, _ := strconv.Atoi(commands.Next())
			x, _ := strconv.Atoi(commands.Next())
			priors[k].insert(x)
		case "extract-min":
			k, _ := strconv.Atoi(commands.Next())
			result.Push(priors[k].getMin())
		case "decrease-key":
			k, _ := strconv.Atoi(commands.Next())
			x, _ := strconv.Atoi(commands.Next())
			y, _ := strconv.Atoi(commands.Next())
			priors[k].decreaseKey(x, y)
		case "merge":
			k, _ := strconv.Atoi(commands.Next())
			m, _ := strconv.Atoi(commands.Next())
			kq, mq := priors[k], priors[m]
			merged := new(PriorityMinHeap)
			var arrmerged [100]int

			j := 0
			for i := 0; i < kq.HeapSize; i++ {
				arrmerged[j] = kq.a[i]
				//merged.insert(kq.a[i])
				j++
			}
			for i := 0; i < mq.HeapSize; i++ {
				arrmerged[j] = mq.a[i]
				//merged.insert(mq.a[i])
				j++
			}
			//merged.build(arrmerged, j)
			merged.a = arrmerged
			merged.HeapSize = j
			priors[lastprior] = merged
			lastprior++
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
		commands.Push(cmd)
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
