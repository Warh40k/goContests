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

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			if table[i][j] != 0 && !hasBorders(i, j, N, table) {
				count++
				break
			}
		}
	}
	return count, table
}

func hasBorders(start, end, N int, table [][]int) bool {
	for i := start; i != end; i++ {
		for j := i + 1; j < N+1; j++ {
			if j == N {
				j = 0
			}
			if end-start != 1 && table[i][j] != 0 {
				return true
			}
			if j == end {
				break
			}

		}
		if i+1 == N {
			i = -1
		}
	}
	return false
}
