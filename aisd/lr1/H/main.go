package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var N int
	fmt.Fscan(in, &N)

	var X = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &X[i])
	}
	count, borders := countBorders(N, X)
	fmt.Fprintf(out, "%d\n%v", count, borders)
	//fmt.Fprintln(out, count)
	out.Flush()
}

func countBorders(N int, X []int) (int, []int) {
	count := 0
	var borders []int
	for i := 1; i < N; i++ {
		for j := 0; j < N; j++ {
			var next = j + i
			if next > N-1 {
				next -= N
			}
			if isConnected(j, i, N, borders) != isConnectedBack(j, next, N, borders) {
				fmt.Println(j, i, N, borders, isConnected(j, i, N, borders))
				fmt.Println(j, next, N, borders, isConnectedBack(j, next, N, borders))
				fmt.Println()
			}
			if isConnected(j, i, N, borders) && X[j]&X[next] != 0 {
				borders = append(borders, next)
				count++
			}
		}

	}
	return count, borders
}
func isConnected(start, step, N int, borders []int) bool {
	if borders != nil {
		for _, val := range borders {
			if (start+step <= N-1 && val > start && val <= start+step) || (start+step > N-1 && (val > start || val < start+step-N)) {
				return false
			}
		}
	}
	return true
}

func isConnectedBack(start, end, N int, borders []int) bool {
	if borders == nil {
		return true
	}
	for i := start + 1; i != end+1; i++ {
		for _, border := range borders {
			if border == i {
				return false
			}
		}
		if i == N {
			i = -1
		}
	}
	return true
}

//for _, border := range borders {
//	if (start < border && end >= border) || (start < border && end >= border) {
//		return false
//	} 1 2 3 4e 5s 6 7
//} (4 + 1 + 1) - 7
// (6-3)
//return true

//func search(x int, X []int) bool {
//	for _, val := range X {
//		if val == x {
//			return true
//		}
//	}
//	return false
//}
