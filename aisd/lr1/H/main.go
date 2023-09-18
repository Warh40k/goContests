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
	out.Flush()
}

func countBorders(N int, X []int) (int, []int) {
	count := 0
	var borders []int
	for i := 0; i < N; i++ {
		var j int
		if i == N-1 {
			j = 0
		} else {
			j = i + 1
		}
		if search(i, X) {
			continue
		}
		for j < N {
			if X[i]&X[j] != 0 {
				borders = append(borders, j-1)
				count++
				if i != N-1 {
					i = j - 1
				}
				break
			}
			j++
		}
	}
	return count, borders
}

func search(el int, X []int) bool {
	for _, val := range X {
		if val == el {
			return true
		}
	}
	return false
}
