package main

import (
	"bufio"
	"fmt"
	"os"
)

var pairs = map[rune][]rune{'X': {'Y', 'Z'}, 'Y': {'Z'}, 'Z': {}}

func checkPair(sym1, sym2 rune) bool {
	parts := pairs[sym1]
	for i := range parts {
		if parts[i] == sym2 {
			return true
		}
	}
	return false
}

func findPair(pos, target, counter, length int, input []rune, used []bool) bool {
	if input[pos] == 'Z' {
		return false
	}
	if !used[pos] && !used[target] && checkPair(input[pos], input[target]) {
		used[pos] = true
		used[target] = true
		counter += 2
		if counter == length {
			return true
		}
		for j := pos + 1; j < length; j++ {
			if !used[j] {
				for k := j + 1; k < length; k++ {
					if !used[k] && findPair(j, k, counter, length, input, used) {
						return true
					}
				}
			}
		}
	}
	used[pos] = false
	used[target] = false

	return false
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var l int
		var sinput string
		fmt.Fscan(in, &l, &sinput)
		input := []rune(sinput)
		found := false
		for j := 1; j < l; j++ {
			used := make([]bool, l)
			if findPair(0, j, 0, l, input, used) {
				found = true
				fmt.Fprintln(out, "Yes")
				break
			}
		}
		if !found {
			fmt.Fprintln(out, "No")
		}
	}
}
