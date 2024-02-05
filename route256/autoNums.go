package main

import (
	"fmt"
	"strconv"
	"strings"
)

var combo1 = "lddll"
var combo2 = "ldll"

func parseString(pattern string, pos int, result []string) bool {
	if pos == len(pattern) {
		fmt.Println(strings.Join(result, " "))
		return true
	} else if pos > len(pattern) {
		return false
	}
	valid1, valid2 := true, true
	var subString1, subString2 string
	if pos+5 <= len(pattern) {
		subString1 = pattern[pos : pos+5]
		for i := range subString1 {
			_, err := strconv.Atoi(string(subString1[i]))
			if (combo1[i] == 'd' && err != nil) ||
				(combo1[i] == 'l' && err == nil) {
				valid1 = false
				break
			}
		}
	} else {
		valid1 = false
	}

	if pos+4 <= len(pattern) {
		subString2 = pattern[pos : pos+4]

		for i := range subString2 {
			_, err := strconv.Atoi(string(subString2[i]))
			if (combo2[i] == 'd' && err != nil) ||
				(combo2[i] == 'l' && err == nil) {
				valid2 = false
				break
			}
		}
	}

	if valid1 {
		new_result := append(result, subString1)
		found := parseString(pattern, pos+5, new_result)
		if found {
			return true
		}
	}
	if valid2 {
		new_result := append(result, subString2)
		found := parseString(pattern, pos+4, new_result)
		if found {
			return true
		}
	}
	return false
}

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var num string
		var result = make([]string, 0)
		fmt.Scan(&num)
		found := parseString(num, 0, result)
		if found == false {
			fmt.Println("-")
		}
	}
}
