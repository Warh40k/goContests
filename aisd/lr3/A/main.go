package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkHeap(n int, values []int) string {
	for i := 1; i <= n; i++ {
		if (2*i <= n && values[i] > values[2*i]) || (2*i+1 <= n && values[i] > values[2*i+1]) {
			return "NO"
		}
	}
	return "YES"
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscan(in, &n)
	values := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &values[i])
	}

	fmt.Fprintln(out, checkHeap(n, values))

	out.Flush()
}
