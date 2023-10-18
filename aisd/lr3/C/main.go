package main

import (
	"bufio"
	"fmt"
	"os"
)

type Worker struct {
	salary int
}

type Value[T any] struct {
	sortValue int
	source    *T
}

type Order struct {
	start, length int
}

type MinHeap[T any] struct {
	a        []*Value[T]
	heapSize int
}

type Queue struct {
	size int
	head *QNode
	tail *QNode
}

type QNode struct {
	value *Order
	next  *QNode
}

func (q *Queue) push(order *Order) {
	top := new(QNode)
	if top == nil {
		panic("StackOverflow")
	}

	top.value = order
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

func (q *Queue) pop() *Order {
	if q.size == 0 {
		panic("Underflow")
	}
	val := q.head.value
	q.head = q.head.next
	q.size--

	return val
}

func (q *Queue) peek() *Order {
	if q.size == 0 {
		panic("Underflow")
	}
	return q.head.value
}

func (bh *MinHeap[T]) sort() []*Value[T] {
	k := bh.heapSize
	var sorted []*Value[T]
	for i := 0; i < k; i++ {
		sorted[i] = bh.getMin()
	}
	return sorted
}

func (bh *MinHeap[T]) siftUp(i int) {
	for bh.a[i].sortValue < bh.a[(i-1)/2].sortValue {
		bh.a[i], bh.a[(i-1)/2] = bh.a[(i-1)/2], bh.a[i]
		i = (i - 1) / 2
	}
}

func (bh *MinHeap[T]) siftDown(i int) {
	for 2*i+1 < bh.heapSize {
		left, right := 2*i+1, 2*i+2
		j := left
		if right < bh.heapSize && bh.a[right].sortValue < bh.a[left].sortValue {
			j = right
		}
		if bh.a[i].sortValue <= bh.a[j].sortValue {
			break
		}
		bh.a[i], bh.a[j] = bh.a[j], bh.a[i]
		i = j
	}
}

func (bh *MinHeap[T]) getMin() *Value[T] {
	if bh.heapSize == 0 {
		panic("Underflow insert")
	}
	hmin := bh.a[0]
	bh.a[0] = bh.a[bh.heapSize-1]
	bh.heapSize--
	bh.siftDown(0)

	return hmin
}

func (bh *MinHeap[T]) insert(key *Value[T]) {
	bh.heapSize += 1
	bh.a[bh.heapSize-1] = key
	bh.siftUp(bh.heapSize - 1)
}

func (bh *MinHeap[T]) build(arr []*Value[T], N int) {
	bh.a = arr
	bh.heapSize = N
	for i := bh.heapSize / 2; i >= 0; i-- {
		bh.siftDown(i)
	}
}

func (bh *MinHeap[T]) peek() *Value[T] {
	if bh.heapSize == 0 {
		return nil
	}
	return bh.a[0]
}

func evaluateSalary(n, m int, incomingOrder *MinHeap[Order], vacantWorkers *MinHeap[Worker]) int {
	workingOrders := &MinHeap[Worker]{a: make([]*Value[Worker], m), heapSize: 0}
	var cost int

	for incomingOrder.heapSize != 0 || workingOrders.heapSize != 0 {
		incomOrd := incomingOrder.peek()
		workOrd := workingOrders.peek()
		if workOrd != nil && (incomingOrder.heapSize == 0 || incomOrd.source.start >= workOrd.sortValue) {
			workingOrders.getMin()
			vacantWorkers.insert(&Value[Worker]{
				source:    workOrd.source,
				sortValue: workOrd.source.salary,
			})
		} else {
			incomingOrder.getMin()
			if vacantWorkers.heapSize != 0 {
				// Достаем самого дешевого шавермана
				worker := vacantWorkers.getMin()
				// Назначение работника на поступивший заказ
				workingOrders.insert(&Value[Worker]{
					source:    worker.source,
					sortValue: incomOrd.source.start + incomOrd.source.length,
				})
				cost += incomOrd.source.length * worker.source.salary
			}
		}
	}
	return cost
}

func main() {
	out := bufio.NewWriter(os.Stdout)
	var n, m int
	fmt.Scan(&n, &m)
	shaurmen := make([]*Value[Worker], n)
	orders := make([]*Value[Order], m)
	vacantWorkers := new(MinHeap[Worker])
	incomingOrders := new(MinHeap[Order])

	for i := 0; i < n; i++ {
		var salary int
		fmt.Scan(&salary)
		shaurmen[i] = &Value[Worker]{
			salary,
			&Worker{salary: salary},
		}
	}

	for i := 0; i < m; i++ {
		var t, f int
		fmt.Scan(&t, &f)
		orders[i] = &Value[Order]{
			t,
			&Order{start: t, length: f},
		}
	}

	vacantWorkers.build(shaurmen, n)
	incomingOrders.build(orders, m)

	fmt.Fprintln(out, evaluateSalary(n, m, incomingOrders, vacantWorkers))
	out.Flush()
}
