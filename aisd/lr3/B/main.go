package main

import (
	"bufio"
	"fmt"
	"goContests/aisd/lr3/B/qu"
	"io"
	"os"
	"strconv"
)

type Queue[T any] struct {
	size     int64
	head     *QNode[T]
	tail     *QNode[T]
	iterator *QNode[T]
}

type QNode[T any] struct {
	value T
	next  *QNode[T]
}

func (q *Queue[T]) resetIterator() {
	q.iterator = q.head
}

func (q *Queue[T]) next() T {
	res := q.iterator
	q.iterator = q.iterator.next
	return res.value
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

type MinHeap struct {
	head      *HeapNode
	tail      *HeapNode
	curparent *HeapNode
}

type HeapNode struct {
	index  int64
	value  int64
	left   *HeapNode
	right  *HeapNode
	next   *HeapNode
	prev   *HeapNode
	parent *HeapNode
}

func (bh *MinHeap) siftUp(node *HeapNode) {
	for node.parent != nil && node.value < node.parent.value {
		node.value, node.parent.value = node.parent.value, node.value
		node = node.parent
	}
}

func (bh *MinHeap) siftDown(node *HeapNode) {
	//var swapee *HeapNode
	//if node == nil {
	//	return
	//}
	//if node.left == nil || node.right.value < node.left.value{
	//	swapee = node.right
	//} else if node.right == nil || node.left.value < node.right.value {
	//	swapee = node.left
	//}
	//if swapee.value < node.value {
	//	node.value, swapee.value = swapee.value, node.value
	//}

	for node.left != nil || node.right != nil {
		var swapee *HeapNode
		if (node.right == nil) || (node.left != nil && node.left.value <= node.right.value) {
			swapee = node.left
		} else if (node.left == nil) || (node.right != nil && node.right.value < node.left.value) {
			swapee = node.right
		}

		if swapee.value < node.value {
			node.value, swapee.value = swapee.value, node.value
			node = swapee
		} else {
			break
		}
	}
}

func (bh *MinHeap) insert(i int64) {
	node := new(HeapNode)
	node.value = i

	if bh.head == nil {
		bh.head = node
		bh.tail = node
		bh.curparent = node
		node.index = 1
		return
	}
	node.index = bh.tail.index + 1
	position := node.index % 2

	if position == 0 {
		node.parent = bh.curparent
		bh.curparent.left = node
	} else {
		node.parent = bh.tail.parent
		bh.tail.parent.right = node
		if bh.curparent.next != nil {
			bh.curparent = bh.curparent.next
		}
	}
	bh.tail.next = node
	node.prev = bh.tail
	bh.tail = node

	bh.siftUp(node)
}

func (bh *MinHeap) getMin() string {
	if bh.tail == nil {
		return "*"
	}
	if bh.head == bh.tail {
		val := strconv.FormatInt(bh.head.value, 10)
		bh.head = nil
		bh.tail = nil

		return val
	}
	hmax := bh.head.value
	bh.head.value = bh.tail.value
	if bh.tail.index%2 == 0 {
		bh.tail.parent.left = nil
		if bh.curparent.prev != nil {
			bh.curparent = bh.curparent.prev
		}
	} else {
		bh.tail.parent.right = nil
	}
	bh.tail = bh.tail.prev
	bh.tail.next = nil
	bh.siftDown(bh.head)

	return strconv.FormatInt(hmax, 10)
}

func (bh *MinHeap) decreaseKey(x, y int64) {
	el := bh.head

	for el != nil {
		if el.value == x {
			el.value = y
			bh.siftUp(el)
			break
		}
		el = el.next
	}
}

func (priors *Queue[T]) findPriority(k int64) T {
	var qu = priors.head
	for i := int64(0); i < k; i++ {
		qu = qu.next
	}
	return qu.value
}

func executeCommands(commands *qu.Queue[string]) *qu.Queue[string] {
	priors := new(qu.Queue[*MinHeap])
	result := new(qu.Queue[string])
	commands.ResetIterator()

	for commands.Iterator != nil {

		switch commands.Next() {
		case "create":
			priors.Push(new(MinHeap))
		case "insert":
			k, _ := strconv.ParseInt(commands.Next(), 10, 64)
			x, _ := strconv.ParseInt(commands.Next(), 10, 64)
			priors.FindPriority(k).insert(x)
		case "extract-min":
			k, _ := strconv.ParseInt(commands.Next(), 10, 64)
			result.Push(priors.FindPriority(k).getMin())
		case "decrease-key":
			k, _ := strconv.ParseInt(commands.Next(), 10, 64)
			x, _ := strconv.ParseInt(commands.Next(), 10, 64)
			y, _ := strconv.ParseInt(commands.Next(), 10, 64)
			priors.FindPriority(k).decreaseKey(x, y)
		case "merge":
			k, _ := strconv.ParseInt(commands.Next(), 10, 64)
			m, _ := strconv.ParseInt(commands.Next(), 10, 64)
			kq, mq := priors.FindPriority(k).head, priors.FindPriority(m).head
			merged := new(MinHeap)
			for kq != nil {
				merged.insert(kq.value)
				kq = kq.next
			}
			for mq != nil {
				merged.insert(mq.value)
				mq = mq.next
			}
			priors.Push(merged)
		}
	}
	return result
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	commands := new(qu.Queue[string])

	for {
		var cmd string
		_, err := fmt.Fscan(in, &cmd)
		if err == io.EOF {
			break
		}
		commands.Push(cmd)
	}

	result := executeCommands(commands).Head
	for result != nil {
		fmt.Fprintln(out, result.Value)
		result = result.Next
	}

	out.Flush()
}
