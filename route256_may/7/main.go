package main

import (
	"bufio"
	"fmt"
	"os"
)

type Car struct {
	Id       int
	Start    int
	End      int
	Capacity int
}

type BhCars struct {
	array []*Car
}

func (h *BhCars) Insert(value *Car) {
	h.array = append(h.array, value)
	h.bubbleUp(len(h.array) - 1)
}

func (h *BhCars) bubbleUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.array[index].Start >= h.array[parentIndex].Start {
			break
		}
		h.array[index], h.array[parentIndex] = h.array[parentIndex], h.array[index]
		index = parentIndex
	}
}

func (h *BhCars) ExtractMin() *Car {
	minim := h.array[0]
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]
	h.sinkDown(0)

	return minim
}

func (h *BhCars) sinkDown(index int) {
	for {
		leftChildIdx := 2*index + 1
		rightChildIdx := 2*index + 2
		minim := index

		if leftChildIdx < len(h.array) && h.array[leftChildIdx].Start < h.array[minim].Start {
			minim = leftChildIdx
		}
		if rightChildIdx < len(h.array) && h.array[rightChildIdx].Start < h.array[minim].Start {
			minim = rightChildIdx
		}

		if minim == index {
			break
		}

		h.array[index], h.array[minim] = h.array[minim], h.array[index]
		index = minim
	}
}

type BhOrders struct {
	array []int
}

func (h *BhOrders) Insert(value int) {
	h.array = append(h.array, value)
	h.bubbleUp(len(h.array) - 1)
}

func (h *BhOrders) bubbleUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.array[index] >= h.array[parentIndex] {
			break
		}
		h.array[index], h.array[parentIndex] = h.array[parentIndex], h.array[index]
		index = parentIndex
	}
}

func (h *BhOrders) ExtractMin() int {
	minim := h.array[0]
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]
	h.sinkDown(0)
	return minim
}

func (h *BhOrders) sinkDown(index int) {
	for {
		leftChildIdx := 2*index + 1
		rightChildIdx := 2*index + 2
		minim := index

		if leftChildIdx < len(h.array) && h.array[leftChildIdx] < h.array[minim] {
			minim = leftChildIdx
		}
		if rightChildIdx < len(h.array) && h.array[rightChildIdx] < h.array[minim] {
			minim = rightChildIdx
		}

		if minim == index {
			break
		}

		h.array[index], h.array[minim] = h.array[minim], h.array[index]
		index = minim
	}
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscan(in, &n)
		orders := &BhOrders{array: make([]int, 0, n)}
		for j := 0; j < n; j++ {
			var arrival int
			fmt.Fscan(in, &arrival)
			orders.Insert(arrival)
		}
		fmt.Fscan(in, &m)
		var cars = &BhCars{array: make([]*Car, 0, m)}

		for j := 0; j < m; j++ {
			car := &Car{Id: j + 1}
			fmt.Fscan(in, &car.Start, &car.End, &car.Capacity)
			cars.Insert(car)
		}
		fmt.Fprintln(out, getCarsForOrders(n, m, orders, cars))
	}
}

func getCarsForOrders(n, m int, orders *BhOrders, cars *BhCars) []int {
	var res = make([]int, n)
	counter := 0
	for i := 0; i < m; i++ {
		car := cars.ExtractMin()
		for j := counter; j < n; j++ {
			order := orders.ExtractMin()
			counter++
			if order >= car.Start && order <= car.End && car.Capacity > 0 {
				res[j] = car.Id
				car.Capacity--
			} else if order > car.Start {
				orders.Insert(order)
				break
			} else {
				res[j] = -1
			}
		}
	}
	return res
}
