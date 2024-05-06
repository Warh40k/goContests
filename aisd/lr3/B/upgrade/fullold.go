package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Vector[T any] struct {
	arr       []T
	size, cap int
}

func (v *Vector[T]) build(cap int) *Vector[T] {
	v.arr = make([]T, cap)
	v.cap = cap

	return v
}

func (v *Vector[T]) append(val T) {
	if v.size == v.cap {
		temp := make([]T, v.cap*2)
		for i := 0; i < v.size; i++ {
			temp[i] = v.arr[i]
		}
		v.arr = temp
		v.cap = v.cap * 2
	}
	v.arr[v.size] = val
	v.size++
}

type MinHeap struct {
	HeapSize int
	A        *Vector[int]
	Heaped   bool
}

func (bh *MinHeap) siftUp(i int) {
	for bh.A.arr[i] < bh.A.arr[(i-1)/2] {
		bh.A.arr[i], bh.A.arr[(i-1)/2] = bh.A.arr[(i-1)/2], bh.A.arr[i]
		i = (i - 1) / 2
	}
}

func (bh *MinHeap) siftDown(i int) {
	for 2*i+1 < bh.HeapSize {
		left, right := 2*i+1, 2*i+2
		j := left
		if right < bh.HeapSize && bh.A.arr[right] < bh.A.arr[left] {
			j = right
		}
		if bh.A.arr[i] <= bh.A.arr[j] {
			break
		}
		bh.A.arr[i], bh.A.arr[j] = bh.A.arr[j], bh.A.arr[i]
		i = j
	}
}

func (bh *MinHeap) insert(i int) {
	bh.HeapSize++
	bh.A.append(i)
	if bh.Heaped {
		bh.siftUp(bh.HeapSize - 1)
	}
}

func (bh *MinHeap) getMin() string {
	if bh.HeapSize == 0 {
		return "*"
	}
	if !bh.Heaped {
		bh.build(bh.A, bh.HeapSize)
	}
	hmax := bh.A.arr[0]
	bh.A.arr[0] = bh.A.arr[bh.HeapSize-1]
	bh.HeapSize--
	bh.A.size--
	bh.siftDown(0)

	return strconv.Itoa(hmax)
}

func (bh *MinHeap) decreaseKey(x, y int) {
	for i := 0; i < bh.HeapSize; i++ {
		if bh.A.arr[i] == x {
			bh.A.arr[i] = y
			if bh.Heaped {
				bh.siftUp(i)
			}
			break
		}
	}
}

func (bh *MinHeap) build(arr *Vector[int], N int) *MinHeap {
	bh.HeapSize = N
	bh.A = arr

	for i := bh.HeapSize / 2; i >= 0; i-- {
		bh.siftDown(i)
	}
	bh.Heaped = true

	return bh
}

func ExecuteCommandsOlder(in *bufio.Reader, out *bufio.Writer) {
	bufsize := 20
	var cmd string
	var k, x, y, m int
	var kq, mq *MinHeap
	var merged *Vector[int]
	priors := new(Vector[*MinHeap]).build(10e5)
	for {
		_, err := fmt.Fscan(in, &cmd)
		if err == io.EOF {
			break
		}
		switch cmd {
		case "create":
			priors.append(new(MinHeap).build(new(Vector[int]).build(bufsize), 0))
		case "insert":
			fmt.Fscan(in, &k, &x)
			priors.arr[k].insert(x)
		case "extract-min":
			fmt.Fscan(in, &k)
			fmt.Fprintln(out, priors.arr[k].getMin())
		case "decrease-key":
			fmt.Fscan(in, &k, &x, &y)
			priors.arr[k].decreaseKey(x, y)
		case "merge":
			fmt.Fscan(in, &k, &m)
			kq, mq = priors.arr[k], priors.arr[m]
			merged = new(Vector[int]).build(bufsize)
			for d := 0; d < kq.HeapSize; d++ {
				merged.append(kq.A.arr[d])
			}
			for d := 0; d < mq.HeapSize; d++ {
				merged.append(mq.A.arr[d])
			}
			priors.append(new(MinHeap).build(merged, merged.size))
		}
	}
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

	ExecuteCommandsOlder(in, out)

	out.Flush()
}
