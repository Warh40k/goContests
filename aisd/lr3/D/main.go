package main

import (
	"bufio"
	"fmt"
	"os"
)

type City struct {
	code, value int
}

type MaxHeap struct {
	a        []*City
	heapSize int
}

func (bh *MaxHeap) siftUp(i int) {
	for bh.a[i].value > bh.a[(i-1)/2].value {
		bh.a[i], bh.a[(i-1)/2] = bh.a[(i-1)/2], bh.a[i]
		i = (i - 1) / 2
	}
}

func (bh *MaxHeap) siftDown(i int) {
	for 2*i+1 < bh.heapSize {
		left, right := 2*i+1, 2*i+2
		j := left
		if right < bh.heapSize && bh.a[right].value > bh.a[left].value {
			j = right
		}
		if bh.a[i].value >= bh.a[j].value {
			break
		}
		bh.a[i], bh.a[j] = bh.a[j], bh.a[i]
		i = j
	}
}

func (bh *MaxHeap) insert(key *City) {
	bh.heapSize += 1
	bh.a[bh.heapSize-1] = key
	bh.siftUp(bh.heapSize - 1)
}

func (bh *MaxHeap) getMax() *City {
	hmax := bh.a[0]
	bh.a[0] = bh.a[bh.heapSize-1]
	bh.heapSize--
	bh.siftDown(0)

	return hmax
}

func (bh *MaxHeap) build(arr []*City, N int) {
	bh.a = arr
	bh.heapSize = N
	for i := bh.heapSize / 2; i >= 0; i-- {
		bh.siftDown(i)
	}
}

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

func (q *Queue[T]) isEmpty() bool {
	if q.size == 0 {
		return true
	}
	return false
}

func (q *Queue[T]) peek() T {
	return q.head.value
}

func checkBanList(search int, banList *Queue[int]) bool {
	item := banList.head

	for item != nil {
		if item.value == search {
			return true
		}
		item = item.next
	}
	return false
}

func findTourPath(n int, answers *Queue[string], rewards, ratings []*City) (*Queue[int], *Queue[int]) {
	ratHeap, rewHeap := new(MaxHeap), new(MaxHeap)
	ratHeap.build(ratings, n)
	rewHeap.build(rewards, n)
	askSeq := new(Queue[int])
	tourSeq := new(Queue[int])
	banList := new(Queue[int])

	for i := 0; i < n; i++ {
		mrat, mrew := ratHeap.getMax(), rewHeap.a[0]
		for checkBanList(mrew.code, banList) {
			rewHeap.getMax()
			mrew = rewHeap.a[0]
		}
		if mrat.code == mrew.code {
			rewHeap.getMax()
			tourSeq.push(mrat.code)
		} else if !answers.isEmpty() {
			if answers.peek() == "YES" {
				tourSeq.push(mrat.code)
			} else {
				banList.push(mrat.code)
			}
			askSeq.push(mrat.code)
			answers.pop()
		}
	}
	return askSeq, tourSeq
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscanln(in, &n)
	rewards, ratings := make([]*City, n), make([]*City, n)

	for i := 0; i < n; i++ {
		var code, rating, reward int
		fmt.Fscanf(in, "%d %d %d\n", &code, &rating, &reward)
		cityRating, cityReward := new(City), new(City)
		cityRating.code = code
		cityRating.value = rating
		cityReward.code = code
		cityReward.value = reward
		rewards[i] = cityReward
		ratings[i] = cityRating
	}
	var m int
	fmt.Fscanln(in, &m)
	answers := new(Queue[string])
	for i := 0; i < m; i++ {
		var ans string
		fmt.Fscan(in, &ans)
		answers.push(ans)
	}

	askSeq, tourSeq := findTourPath(n, answers, rewards, ratings)

	for askSeq.size != 0 {
		fmt.Fprint(out, askSeq.pop(), " ")
	}
	fmt.Fprintln(out)
	for tourSeq.size != 0 {
		fmt.Fprint(out, tourSeq.pop(), " ")
	}

	out.Flush()
}
