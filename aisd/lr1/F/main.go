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

	X := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &X[i])
	}
	fmt.Fprintln(out, sleepDay(N, X))
	out.Flush()
}

func sleepDay(N int, X []int) int {
	sforward, sbackward := make([]int, N), make([]int, N)
	sforward[0] = X[0]
	sbackward[N-1] = X[N-1]

	for i := 1; i < N; i++ {
		sforward[i] = sforward[i-1] + X[i]
	}

	for i := N - 2; i > 0; i-- {
		sbackward[i] = sbackward[i+1] + X[i]
	}

	for i := 0; i < N; i++ {
		if sforward[i] == sbackward[i] {
			return i
		}
	}
	return -1
}
