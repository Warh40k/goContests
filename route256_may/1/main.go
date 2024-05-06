package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b int
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fscan(in, &a, &b)
	fmt.Fprintln(out, a-b)
}
