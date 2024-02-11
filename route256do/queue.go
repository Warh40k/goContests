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
	for i := pos + 1; i < len(input); i++ {

		if !used[pos] && !used[i] && checkPair(input[pos], input[i]) {
			used[pos] = true
			used[i] = true
			counter += 2
			if counter == length {
				return true
			}
			for j := pos + 1; j < len(input); j++ {
				if used[j] == false {
					if findPair(j, counter, length, input, used) {
						return true
					}
				}
			}
		}
	}
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
		fmt.Fprintln(out, input)
		used := make([]bool, l)
		if findPair(0, 0, l, input, used) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
