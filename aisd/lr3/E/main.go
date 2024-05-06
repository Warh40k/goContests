package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type MinHeap struct {
	a        []int
	heapSize int
}

func (bh *MinHeap) sort(out *bufio.Writer) {
	k := bh.heapSize
	for i := 0; i < k; i++ {
		fmt.Fprint(out, bh.getMin(), " ")
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

func (bh *MinHeap) getMin() int {
	hmax := bh.a[0]
	bh.a[0] = bh.a[bh.heapSize-1]
	bh.heapSize--
	bh.siftDown(0)

	return hmax
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscan(in, &n)
	values := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &values[i])
	}

	bh := new(MinHeap)
	bh.a = values
	bh.heapSize = n
	for i := n / 2; i >= 0; i-- {
		bh.siftDown(i)
	}

	bh.sort(out)

	out.Flush()
}

func test(out *bufio.Writer) {
	for {
		n := int(10e5)
		values := make([]int, n)

		for i := 0; i < n; i++ {
			values[i] = rand.Intn(10e9-1) + 1
		}
		bh := new(MinHeap)
		bh.heapSize = n
		bh.a = values
		bh.sort(out)
		out.Flush()
	}
}
