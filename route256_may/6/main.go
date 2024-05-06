package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n, k, m int
		fmt.Fscan(in, &n, &k, &m)
		var boxes = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &boxes[j])
		}
		fmt.Fprintln(out, minPaths(n, k, m, boxes))
	}
}

func minPaths(n, k, m int, boxes []int) int {
	var cars = make([]int, n)
	var curCar, numRoads int
	numRoads = 1

	for i := m - 1; i >= 0; i-- {
		weight := 1 << boxes[i]
		if cars[curCar]+weight <= k {
			cars[curCar] += weight
		} else if curCar < n-1 {
			curCar++
			cars[curCar] = weight
		} else {
			numRoads++
			curCar = 0
			cars[curCar] = weight
		}
	}
	return numRoads
}
