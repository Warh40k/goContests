package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var n, k int
	fmt.Fscan(in, &n, &k)
	arr := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	fmt.Fprintln(out, strings.Join(shift(arr, k, n), " "))
	out.Flush()
}

func shift(arr []string, k int, n int) []string {
	result := make([]string, n)
	index := int(math.Abs(float64(k % n)))

	if k > 0 {
		index = n - index
	}

	for i, j := 0, index; i < n; i, j = i+1, j+1 {
		if j == n {
			j = 0
		}
		result[i] = arr[j]
	}

	return result
}
