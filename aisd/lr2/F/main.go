package main

import (
	"bufio"
	"fmt"
	"os"
)

type MaxHeap struct {
	a                 []int
	maxSize, heapSize int
}

type MinHeap struct {
	a                 []int
	maxSize, heapSize int
}

func (bh *MaxHeap) siftUp(i int) {
	for bh.a[i] > bh.a[(i-1)/2] {
		bh.a[i], bh.a[(i-1)/2] = bh.a[(i-1)/2], bh.a[i]
		i = (i - 1) / 2
	}
}

func (bh *MaxHeap) siftDown(i int) {
	for 2*i+1 < bh.heapSize {
		left, right := 2*i+1, 2*i+2
		j := left
		if right < bh.heapSize && bh.a[right] > bh.a[left] {
			j = right
		}
		if bh.a[i] >= bh.a[j] {
			break
		}
		bh.a[i], bh.a[j] = bh.a[j], bh.a[i]
		i = j
	}
}

func (bh *MaxHeap) insert(key int) {
	bh.heapSize += 1
	bh.a[bh.heapSize-1] = key
	bh.siftUp(bh.heapSize - 1)
}

func (bh *MaxHeap) getMax() int {
	hmax := bh.a[0]
	bh.a[0] = bh.a[bh.heapSize-1]
	bh.heapSize--
	bh.siftDown(0)

	return hmax
}

func (bh *MaxHeap) build(arr []int, N int) {
	bh.a = arr
	bh.heapSize = N
	bh.maxSize = N
	for i := bh.heapSize / 2; i >= 0; i-- {
		bh.siftDown(i)
	}
}

func (bh *MinHeap) siftUp(i int) {
	for bh.a[i] < bh.a[(i-1)/2] {
		bh.a[i], bh.a[(i-1)/2] = bh.a[(i-1)/2], bh.a[i]
		i = (i - 1) / 2
	}
}

func (bh *MinHeap) siftDown(i int) {
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

func (bh *MinHeap) insert(key int) {
	bh.heapSize += 1
	bh.a[bh.heapSize-1] = key
	bh.siftUp(bh.heapSize - 1)
}

func (bh *MinHeap) getMin() int {
	hmax := bh.a[0]
	bh.a[0] = bh.a[bh.heapSize-1]
	bh.heapSize--
	bh.siftDown(0)

	return hmax
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var k, n int
	var ratios = make([][]int, 2)
	fmt.Fscan(in, &n, &k)

	for i := range ratios {
		ratios[i] = make([]int, n)
		for j := range ratios[i] {
			fmt.Fscan(in, &ratios[i][j])
		}
	}

	fmt.Fprintln(out, getKRatio(k-1, n, ratios))

	out.Flush()
}

func getKRatio(k, n int, ratios [][]int) int {
	t := n*n - k

	if t > k {
		return maxHeap(k, n, ratios)
	} else {
		return minHeap(t+1, n, ratios)
	}

}

func maxHeap(k, n int, ratios [][]int) int {
	htree := new(MaxHeap)
	htree.a = make([]int, k)
	var sum int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum = ratios[0][i] + ratios[1][j]
			if htree.heapSize < k {
				htree.insert(sum)
			} else if sum < htree.a[0] {
				htree.getMax()
				htree.insert(sum)
			}
		}
	}

	return htree.getMax()
}

func minHeap(k, n int, ratios [][]int) int {
	htree := new(MinHeap)
	htree.a = make([]int, k)
	var sum int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum = ratios[0][i] + ratios[1][j]
			if htree.heapSize < k {
				htree.insert(sum)
			} else if sum > htree.a[0] {
				htree.getMin()
				htree.insert(sum)
			}
		}
	}

	return htree.getMin()
}
