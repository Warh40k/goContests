package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type Worker struct {
	salary, timeEnd int
}

type Order struct {
	start, length int
}

type SalaryMinHeap struct {
	a        []*Worker
	heapSize int
}

type TimeEndMinHeap struct {
	a        []*Worker
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
		return nil
	}
	return q.head.value
}

func (bh *SalaryMinHeap) siftDown(i int) {
	for bh.a[i].salary < bh.a[(i-1)/2].salary {
		bh.a[i], bh.a[(i-1)/2] = bh.a[(i-1)/2], bh.a[i]
		i = (i - 1) / 2
	}
}

func (bh *SalaryMinHeap) siftUp(i int) {
	for 2*i+1 < bh.heapSize {
		left, right := 2*i+1, 2*i+2
		j := left
		if right < bh.heapSize && bh.a[right].salary < bh.a[left].salary {
			j = right
		}
		if bh.a[i].salary <= bh.a[j].salary {
			break
		}
		bh.a[i], bh.a[j] = bh.a[j], bh.a[i]
		i = j
	}
}

func (bh *SalaryMinHeap) getMin() *Worker {
	if bh.heapSize == 0 {
		panic("Underflow insert")
	}
	hmin := bh.a[0]
	bh.a[0] = bh.a[bh.heapSize-1]
	bh.heapSize--
	bh.siftUp(0)

	return hmin
}

func (bh *SalaryMinHeap) insert(key *Worker) {
	bh.heapSize += 1
	bh.a[bh.heapSize-1] = key
	bh.siftDown(bh.heapSize - 1)
}

func (bh *SalaryMinHeap) build(arr []*Worker, N int) {
	bh.a = arr
	bh.heapSize = N
	for i := bh.heapSize / 2; i >= 0; i-- {
		bh.siftUp(i)
	}
}

func (bh *SalaryMinHeap) peek() *Worker {
	if bh.heapSize == 0 {
		return nil
	}
	return bh.a[0]
}

func (bh *TimeEndMinHeap) siftDown(i int) {
	for bh.a[i].timeEnd < bh.a[(i-1)/2].timeEnd {
		bh.a[i], bh.a[(i-1)/2] = bh.a[(i-1)/2], bh.a[i]
		i = (i - 1) / 2
	}
}

func (bh *TimeEndMinHeap) siftUp(i int) {
	for 2*i+1 < bh.heapSize {
		left, right := 2*i+1, 2*i+2
		j := left
		if right < bh.heapSize && bh.a[right].timeEnd < bh.a[left].timeEnd {
			j = right
		}
		if bh.a[i].timeEnd <= bh.a[j].timeEnd {
			break
		}
		bh.a[i], bh.a[j] = bh.a[j], bh.a[i]
		i = j
	}
}

func (bh *TimeEndMinHeap) getMin() *Worker {
	if bh.heapSize == 0 {
		panic("Underflow insert")
	}
	hmin := bh.a[0]
	bh.a[0] = bh.a[bh.heapSize-1]
	bh.heapSize--
	bh.siftUp(0)

	return hmin
}

func (bh *TimeEndMinHeap) insert(key *Worker) {
	bh.heapSize += 1
	bh.a[bh.heapSize-1] = key
	bh.siftDown(bh.heapSize - 1)
}

func (bh *TimeEndMinHeap) build(arr []*Worker, N int) {
	bh.a = arr
	bh.heapSize = N
	for i := bh.heapSize / 2; i >= 0; i-- {
		bh.siftUp(i)
	}
}

func (bh *TimeEndMinHeap) peek() *Worker {
	if bh.heapSize == 0 {
		return nil
	}
	return bh.a[0]
}

func evaluateSalary(m int, incomingOrder *Queue, vacantWorkers *SalaryMinHeap) int {
	timeEndHeap := &TimeEndMinHeap{a: make([]*Worker, m), heapSize: 0}
	var cost int
	var incomOrd *Order
	var worker, vacant *Worker

	for incomingOrder.size != 0 {
		incomOrd = incomingOrder.peek()
		worker = timeEndHeap.peek()
		if worker != nil && incomOrd.start >= worker.timeEnd {
			//fmt.Println("Закончить работу")
			timeEndHeap.getMin()
			vacantWorkers.insert(worker)
		} else {
			//fmt.Println("Назначить заказ")
			incomingOrder.pop()
			if vacantWorkers.heapSize != 0 {
				// Достаем самого дешевого шавермана
				vacant = vacantWorkers.getMin()
				vacant.timeEnd = incomOrd.start + incomOrd.length
				// Назначение работника на поступивший заказ
				timeEndHeap.insert(vacant)
				cost += incomOrd.length * vacant.salary
			}
		}
	}
	return cost
}

func test() {
	n, m := int(10e5), int(10e5)
	var t, f, salary int
	var shaurmen = make([]*Worker, n)
	vacantWokers := new(SalaryMinHeap)
	var incomingOrders = new(Queue)
	for i := 0; i < n; i++ {
		salary = rand.Intn(1000) + 1
		shaurmen[i] = &Worker{salary: salary}
	}
	start := rand.Intn(10e4)
	for i := 0; i < m; i++ {
		t = start + rand.Intn(10e3)
		f = t + rand.Intn(10e3) + 1
		incomingOrders.push(&Order{start: t, length: f})
	}
	vacantWokers.build(shaurmen, n)
	fmt.Println(evaluateSalary(m, incomingOrders, vacantWokers))
}

func main() {
	out := bufio.NewWriter(os.Stdout)
	var n, m, salary, t, f int
	fmt.Scan(&n, &m)
	shaurmen := make([]*Worker, n)
	incomingOrders := new(Queue)
	vacantWorkers := new(SalaryMinHeap)

	for i := 0; i < n; i++ {
		fmt.Scan(&salary)
		shaurmen[i] = &Worker{salary: salary}
	}

	for i := 0; i < m; i++ {
		fmt.Scan(&t, &f)
		incomingOrders.push(&Order{start: t, length: f})
	}

	vacantWorkers.build(shaurmen, n)

	fmt.Fprintln(out, evaluateSalary(m, incomingOrders, vacantWorkers))
	out.Flush()
}
