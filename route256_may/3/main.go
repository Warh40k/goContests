package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var s string
		fmt.Fscan(in, &s)
		if checkTransform(s) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func checkTransform(s string) bool {
	n := len(s)
	if s[0] != s[n-1] {
		return false
	}
	chsym := s[0]
	for i := 0; i < n-1; i++ {
		if s[i] != chsym && s[i+1] != chsym {
			return false
		}
	}

	return true
}
