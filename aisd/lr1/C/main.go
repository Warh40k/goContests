package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var a, b int
	fmt.Fscan(in, &a, &b)
	fmt.Println(maxFriends(a, b))
	out.Flush()
}

func maxFriends(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
