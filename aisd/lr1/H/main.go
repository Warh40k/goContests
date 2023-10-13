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
	//test(out)
	//os.Exit(1)
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

func test(writer *bufio.Writer) {
	N := int(4 * math.Pow(10, 5))
	arr := make([]int, N)

	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(int(math.Pow(10, 9)) + 1)
	}
	count, borders := countBorders(N, arr)
	fmt.Println(writer, count, borders)
	writer.Flush()
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
			if X[j]&X[next] != 0 && (i == 1 || hasBordersBetween(j, i, N, borders)) {
				borders = append(borders, next)
				count++
			}
		}
	}
	if count == 0 {
		count = 1
	}
	return count, borders
}
func hasBordersBetween(start, step, N int, borders []int) bool {
	if borders != nil {
		for _, val := range borders {
			if (start+step <= N-1 && val > start && val <= start+step) ||
				(start+step > N-1 && (val > start || val <= start+step-N)) {
				return false
			}
		}
	}
	return true
}
