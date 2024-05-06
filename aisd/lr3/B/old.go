package main

import (
	"goContests/aisd/lr3/B/qu"
	"strconv"
)

type PriorityMinHeap struct {
	k, heapSize int
	a           [20]int
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
	for i := 0; i < bh.heapSize; i++ {
		if bh.a[i] == x {
			bh.a[i] = y
			bh.siftUp(i)
			break
		}
	}
}

func executeCommandsOld(commands *qu.Queue[string]) *qu.Queue[string] {
	priors := new(qu.Queue[*PriorityMinHeap])
	result := new(qu.Queue[string])

	for commands.Size != 0 {

		switch commands.Pop() {
		case "create":
			priors.Push(new(PriorityMinHeap))
		case "insert":
			k, _ := strconv.Atoi(commands.Pop())
			x, _ := strconv.Atoi(commands.Pop())
			priors.FindPriority(int64(k)).insert(x)
		case "extract-min":
			k, _ := strconv.Atoi(commands.Pop())
			result.Push(priors.FindPriority(int64(k)).getMin())
		case "decrease-key":
			k, _ := strconv.Atoi(commands.Pop())
			x, _ := strconv.Atoi(commands.Pop())
			y, _ := strconv.Atoi(commands.Pop())
			priors.FindPriority(int64(k)).decreaseKey(x, y)
		case "merge":
			k, _ := strconv.Atoi(commands.Pop())
			m, _ := strconv.Atoi(commands.Pop())
			kq, mq := priors.FindPriority(int64(k)), priors.FindPriority(int64(m))
			merged := new(PriorityMinHeap)
			for i := 0; i < kq.heapSize; i++ {
				merged.insert(kq.a[i])
			}
			for i := 0; i < mq.heapSize; i++ {
				merged.insert(mq.a[i])
			}
			priors.Push(merged)
		}
	}
	return result
}
