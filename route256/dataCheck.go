package main

import (
	"bufio"
	"fmt"
	"os"
)

func leapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	}
	return false
}

func isBigMonth(month int) bool {
	var bigM = []int{1, 3, 5, 7, 8, 10, 12}
	for i := range bigM {
		if bigM[i] == month {
			return true
		}
	}
	return false
}

func checkDate(day, month, year int) bool {
	var febLen = 28
	isLeap := leapYear(year)
	if isLeap {
		febLen = 29
	}
	if day == 31 && !isBigMonth(month) {
		return false
	} else if month == 2 && day > febLen {
		return false
	}
	return true
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var day, month, year int
		fmt.Fscan(in, &day, &month, &year)
		if checkDate(day, month, year) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
	out.Flush()
}
