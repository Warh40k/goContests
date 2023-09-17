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
	fmt.Fprintln(out, a+b)
	out.Flush()
}
