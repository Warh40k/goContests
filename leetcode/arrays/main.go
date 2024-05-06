package main

import (
	"fmt"
	"slices"
)

func removeDuplicates(nums []int) int {
	l := len(nums)
	j := 0
	count := 1
	for i := 1; i < l; i++ {
		if nums[i] != nums[j] {
			nums = slices.Delete(nums, j, i-1)
			l -= i - j - 1
			j++
			i = j
			count++
		}
	}

	return count
}

func main() {
	var n int
	fmt.Scan(&n)
	var nums = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	count := removeDuplicates(nums)
	fmt.Println(count)
}
