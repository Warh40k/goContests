package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var a, b, c, d int
	fmt.Fscan(in, &a, &b, &c, &d)

	total := a
	diff := d - b
	if diff > 0 {
		total += c * diff
	}
	fmt.Fprintln(out, total)
}
