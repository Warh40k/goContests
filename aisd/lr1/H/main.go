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

	count, _ := countBorders(N, X)
	fmt.Fprintln(out, count)
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
			k := hasBordersBetween(j, i, N, borders)
			if (i == 1 || k == -1) && X[j]&X[next] != 0 {
				borders = append(borders, next)
				count++
			} else if k != -1 && k > j {
				j = k
			}
		}
	}
	if count == 0 {
		count = 1
	}
	return count, borders
}
func hasBordersBetween(start, step, N int, borders []int) int {
	if borders != nil {
		for _, val := range borders {
			if (start+step <= N-1 && val > start && val <= start+step) ||
				(start+step > N-1 && (val > start || val < start+step-N)) {
				return val
			}
		}
	}
	return -1
}
