package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func compress1(adata []int, length int) (int, []int) {
	result := make([]int, 0, length)
	for i := 0; i < length; i++ {
		var asc = true
		var count = 0
		result = append(result, adata[i])
		if i != length-1 && adata[i] == adata[i+1]+1 {
			asc = false
		}
		for j := i + 1; j < length; j++ {
			if asc && adata[j] == adata[j-1]+1 {
				count++
			} else if !asc && adata[j] == adata[j-1]-1 {
				count--
			} else {
				break
			}
		}
		i += int(math.Abs(float64(count)))
		result = append(result, count)
	}

	return len(result), result
}

func compress2(adata []int, length int) (int, []int) {
	result := make([]int, 0, length)
	for i := length - 1; i >= 0; i-- {
		var asc = true
		var count = 0

		if i != 0 && adata[i] != adata[i-1]+1 {
			asc = false
		}
		for j := i; j > 0; j-- {
			if asc && adata[j] == adata[j-1]+1 {
				count++
			} else if !asc && adata[j] == adata[j-1]-1 {
				count--
			} else {
				break
			}
		}
		i -= int(math.Abs(float64(count)))
		result = append(result, count)
		result = append(result, adata[i])
	}

	return len(result), result
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Fscan(in, &k)
		data := make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscan(in, &data[j])
		}
		m1, result1 := compress1(data, k)
		m2, result2 := compress2(data, k)

		if m1 <= m2 {
			fmt.Fprintln(out, m1)
			for j := range result1 {
				fmt.Fprint(out, strconv.Itoa(result1[j])+" ")
			}
			fmt.Fprintln(out)
		} else {
			fmt.Fprintln(out, m2)
			for j := m2 - 1; j >= 0; j-- {
				fmt.Fprint(out, strconv.Itoa(result2[j])+" sfsf")
			}
			fmt.Fprintln(out)
		}
	}
	out.Flush()
}
