package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func countSeasons(n, k int, input []int) int {
	var asc bool
	var maxCount int

	for i := 0; i < n-2*k; i++ {
		var j int
		asc = true
		curCount := 0
		m := 0
		for j = i; j < n; j++ {
			if m == k {
				asc = false
			}
			m++
			if m == 2*k+1 {
				curCount++
				m = 1
				asc = true
			}
			if j != n-1 && ((asc && input[j] >= input[j+1]) || (!asc && input[j] <= input[j+1])) {
				break
			}

		}
		if curCount > maxCount {
			maxCount = curCount
		}
	}
	return maxCount
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var l int
		fmt.Fscan(in, &l)
		var input = make([]int, l)
		var result = make([]int, l)
		for j := 0; j < l; j++ {
			fmt.Fscan(in, &input[j])
		}
		for k := 1; k < l+1; k++ {
			result[k-1] = countSeasons(l, k, input)
			fmt.Fprint(out, strconv.Itoa(result[k-1])+" ")
		}
		fmt.Fprintln(out)
	}
}
