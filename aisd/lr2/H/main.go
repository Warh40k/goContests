package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var n, k int
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &k)
	places := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &places[i])
	}

	fmt.Fprintln(out, evaluateDistance(n, k, places))

	out.Flush()
}

func evaluateDistance(n, k int, places []int) int {
	left := 1
	right := places[n-1] - places[0]
	var result = -1

	for left <= right {
		mid := (right + left) / 2
		temp := places[0]
		count := 1

		for i := 1; i < n; i++ {
			if places[i]-temp >= mid {
				temp = places[i]
				count++
			}
		}

		if count >= k {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}
