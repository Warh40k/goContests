package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var n int
	var N string
	fmt.Fscan(in, &n, &N)
	fmt.Fprintln(out, checkPallindrome(n, N))
	out.Flush()
}

func checkPallindrome(n int, N string) string {
	var left, right, checked = 0, n - 1, false

	for left < right {
		if N[left] == N[right] {
			left++
			right--
		} else if !checked {
			if N[left+1] == N[right] {
				left++
				checked = true
			} else if N[left] == N[right-1] {
				right--
				checked = true
			} else {
				return "NO"
			}
		}
	}

	return "YES"
}
