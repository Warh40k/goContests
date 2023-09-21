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

	count, table := countBorders(N, X)
	fmt.Fprintln(out, count)
	fmt.Fprintln(out, "", X)
	for i, val := range table {
		fmt.Fprint(out, X[i], "")
		fmt.Fprintln(out, val)
	}
	out.Flush()
}

func countBorders(N int, X []int) (int, [][]int) {
	count := 0
	table := make([][]int, N)
	for i := 0; i < N; i++ {
		table[i] = make([]int, N)
		for j := 0; j < N; j++ {
			table[i][j] = X[i] & X[j]
		}
	}
	//for i := 0; i < N; i++ {
	//	table[i] = make([]int, N)
	//	for j := i + 1; j < N; j++ {
	//		if table[j][j] != 0 && !hasBorders(i, j, N, table) {
	//			count++
	//		}
	//	}
	//}
	var i, m, j = 0, 1, 1
	for i != 0 || j < N {
		if table[i][j] != 0 && !hasBorders(i, j, N, table) {
			count++
		}
		if j == N-1 {
			m++
			j = m
			i = 0
		} else {
			i++
			j++
		}
	}
	return count, table
}

func hasBorders(f, s, N int, table [][]int) bool {
	for i := f; i < N-1; i++ {
		for j := i + 1; j <= s; j++ {
			if i == f && j == s {
				continue
			}
			if table[i][j] != 0 {
				return true
			}
		}
	}
	return false
}
