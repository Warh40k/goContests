package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
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
	test(out)
	fmt.Fprintln(out, sleepDayTest(N, X))
	out.Flush()
}

func test(out *bufio.Writer) {
	for {
		//N := rand.Intn(8*int(math.Pow(10, 6))) + 1
		N := 8 * int(math.Pow(10, 6))
		X := make([]int, N)
		for i := range X {
			X[i] = rand.Intn(2000) - 1000 + 1
		}
		//a := sleepDay(N, X)
		b := sleepDayTest(N, X)
		//if a != b {
		//	fmt.Fprintln(out, X, a, b)
		//}
		fmt.Fprintln(out, b)
		out.Flush()
	}

}

func sleepDay(N int, X []int) int {
	var sbackward, sforward = X[0], 0
	if N == 1 {
		return 0
	}
	for i := 1; i < N; i++ {
		for j := i + 1; j < N; j++ {
			sforward += X[j]
			if sforward > sbackward {
				break
			} else if sbackward == sforward {
				return i
			}
		}
		sbackward += X[i]
		sforward = 0
	}
	return -1
}

func sleepDayTest(N int, X []int) int {
	sforward, sbackward := make([]int, N), make([]int, N)
	sforward[0] = X[0]
	sbackward[N-1] = X[N-1]

	for i := 1; i < N; i++ {
		sforward[i] = sforward[i-1] + X[i]
		sbackward[N-i-1] = sbackward[N-i] + X[N-i-1]
	}

	for i := 0; i < N; i++ {
		if sforward[i] == sbackward[i] {
			return i
		}
	}
	return -1
}
