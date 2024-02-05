package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var shipCounts = []int{0, 4, 3, 2, 1}
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscan(in, &n)
	var input = make([][]int, n)
	var ships = make([][]int, n)

	for i := 0; i < n; i++ {
		input[i] = make([]int, 10)
		for j := 0; j < 10; j++ {
			fmt.Fscan(in, &input[i][j])
		}
	}

	for i := 0; i < n; i++ {
		right := true
		ships[i] = make([]int, 5)
		for j := 0; j < 10; j++ {
			shipSize := input[i][j]
			ships[i][shipSize]++
			if ships[i][shipSize] > shipCounts[shipSize] {
				right = false
				fmt.Fprintln(out, "NO")

				break
			}
		}
		if right {
			fmt.Fprintln(out, "YES")
		}
	}

	out.Flush()
}
