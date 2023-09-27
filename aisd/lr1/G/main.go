package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

	var t, n int
	fmt.Fscan(in, &t)
	var votes [][]string

	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		votes = make([][]string, n)
		for j := 0; j < n; j++ {
			votes[j] = make([]string, 2)
			fmt.Fscan(in, &votes[j][0], &votes[j][1])
		}
		findCompomiss(out, votes)
	}

	out.Flush()
}

func findCompomiss(out *bufio.Writer, votes [][]string) {
	var vmin, vmax, val = 15, 30, 0
	for _, vote := range votes {
		val, _ = strconv.Atoi(vote[1])
		if vote[0] == ">=" && vmin < val {
			vmin = val
		} else if vote[0] == "<=" && vmax > val {
			vmax = val
		}
		if vmin > vmax {
			fmt.Fprintln(out, -1)
		} else {
			fmt.Fprintln(out, vmin)
		}
	}
}
