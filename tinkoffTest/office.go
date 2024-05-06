package main

import (
	"bufio"
	"fmt"
	"os"
)

func minPath(n, t, wt int, workerFloors []int) int {
	floorPath := workerFloors[n-1] - workerFloors[0]
	floorT := workerFloors[wt-1]
	pathToTBottom := floorT - workerFloors[0]
	pathToTTop := floorPath - floorT
	var sameFloor = 0
	//for i := 0; i < n-1; i++ {
	//	if workerFloors[i] == workerFloors[i+1] {
	//		sameFloor++
	//	}
	//}

	if pathToTBottom <= t || pathToTTop <= t {
		return floorPath - sameFloor
	} else {
		var extraPath int
		if pathToTBottom < pathToTTop {
			extraPath = pathToTBottom
		} else {
			extraPath = pathToTTop
		}
		return floorPath + extraPath - sameFloor
	}
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, t, wt int
	fmt.Fscan(in, &n, &t)

	var floors = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &floors[i])
	}
	fmt.Fscan(in, &wt)

	fmt.Println(minPath(n, t, wt, floors))
}
