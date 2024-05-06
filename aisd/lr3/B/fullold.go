package main

import (
	"strconv"
)

//type Queue[T any] struct {
//	size     int64
//	head     *QNode[T]
//	tail     *QNode[T]
//	iterator *QNode[T]
//}
//
//type QNode[T any] struct {
//	value T
//	next  *QNode[T]
//}
//
//func (q *Queue[T]) resetIterator() {
//	q.iterator = q.head
//}
//
//func (q *Queue[T]) next() T {
//	res := q.iterator
//	q.iterator = q.iterator.next
//	return res.value
//}
//
//func (q *Queue[T]) push(el T) {
//	top := new(QNode[T])
//	if top == nil {
//		panic("StackOverflow")
//	}
//
//	top.value = el
//	// Если раньше элементов не было - устанавливаем head и tail на top
//	if q.size == 0 {
//		q.head = top
//		q.tail = top
//	} else {
//		// Иначе создаем ссылку предыдущего элемента на текущий и делаем текущий эл. последним
//		q.tail.next = top
//		q.tail = top
//	}
//
//	q.size++
//}
//
//func (q *Queue[T]) pop() T {
//	if q.size == 0 {
//		panic("Underflow")
//	}
//	val := q.head.value
//	q.head = q.head.next
//	q.size--
//
//	return val
//}

type MinHeapOld struct {
	k, heapSize int
	a           [100]int
}

func (bh *MinHeapOld) siftUp(i int) {
	for bh.a[i] < bh.a[(i-1)/2] {
		bh.a[i], bh.a[(i-1)/2] = bh.a[(i-1)/2], bh.a[i]
		i = (i - 1) / 2
	}
}

func (bh *MinHeapOld) siftDown(i int) {
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

func (bh *MinHeapOld) insert(i int) {
	bh.heapSize++
	bh.a[bh.heapSize-1] = i
	bh.siftUp(bh.heapSize - 1)
}

func (bh *MinHeapOld) getMin() string {
	if bh.heapSize == 0 {
		return "*"
	}
	hmax := bh.a[0]
	bh.a[0] = bh.a[bh.heapSize-1]
	bh.heapSize--
	bh.siftDown(0)

	return strconv.Itoa(hmax)
}

func (bh *MinHeapOld) decreaseKey(x, y int) {
	for i := 0; i < bh.heapSize; i++ {
		if bh.a[i] == x {
			bh.a[i] = y
			bh.siftUp(i)
			break
		}
	}
}

func (bh *MinHeapOld) build(arr [100]int, N int) {
	bh.a = arr
	bh.heapSize = N
	for i := bh.heapSize / 2; i >= 0; i-- {
		bh.siftUp(i)
	}
}

func ExecuteCommandsOlder(commands *Queue[string]) *Queue[string] {
	priors := make([]*MinHeapOld, 10e6)
	result := new(Queue[string])
	commands.resetIterator()
	lastprior := 0
	for commands.iterator != nil {

		switch commands.next() {
		case "create":
			priors[lastprior] = new(MinHeapOld)
			lastprior++
		case "insert":
			k, _ := strconv.Atoi(commands.next())
			x, _ := strconv.Atoi(commands.next())
			priors[k].insert(x)
		case "extract-min":
			k, _ := strconv.Atoi(commands.next())
			result.push(priors[k].getMin())
		case "decrease-key":
			k, _ := strconv.Atoi(commands.next())
			x, _ := strconv.Atoi(commands.next())
			y, _ := strconv.Atoi(commands.next())
			priors[k].decreaseKey(x, y)
		case "merge":
			k, _ := strconv.Atoi(commands.next())
			m, _ := strconv.Atoi(commands.next())
			kq, mq := priors[k], priors[m]
			merged := new(MinHeapOld)
			var arrmerged [100]int

			j := 0
			for i := 0; i < kq.heapSize; i++ {
				arrmerged[j] = kq.a[i]
				j++
			}
			for i := 0; i < mq.heapSize; i++ {
				arrmerged[j] = mq.a[i]
				j++
			}
			merged.build(arrmerged, j)
			priors[lastprior] = merged
			lastprior++
		}
	}
	return result
}

//func main() {
//	out := bufio.NewWriter(os.Stdout)
//	commands := new(Queue[string])
//
//	for {
//		var cmd string
//		_, err := fmt.Scan(&cmd)
//		if err == io.EOF {
//			break
//		}
//		commands.push(cmd)
//	}
//
//	result := ExecuteCommandsOlder(commands).head
//	for result != nil {
//		if result.next != nil {
//			fmt.Fprintln(out, result.value)
//		} else {
//			fmt.Fprint(out, result.value)
//		}
//		result = result.next
//	}
//	out.Flush()
//}
