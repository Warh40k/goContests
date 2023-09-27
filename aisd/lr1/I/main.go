package main

import (
	"fmt"
)

func main() {
	var start, end uint64
	fmt.Scan(&start, &end)

	fmt.Println(countTradeDays(start, end))
}

func countTradeDays(start, end uint64) int {
	count := 0
	for i := start; i <= end; i++ {
		if checkBinary(i) {
			count++
		}
	}
	return count
}

func checkBinary(num uint64) bool {
	prev := num % 2
	var cur uint64 = 0
	factor := num / 2
	gcount := 0

	for factor != 0 {
		cur = factor % 2
		if cur != prev {
			gcount++
			prev = cur
		}
		if gcount > 2 {
			return false
		}
		if factor != 0 {
			factor = factor / 2
		}
	}
	if gcount != 2 {
		return false
	}
	return true
}
