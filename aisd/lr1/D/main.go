package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var n, m, q, seas, ep int
	fmt.Fscan(in, &n, &m, &q)

	avail := make([][]int, m)
	for i := range avail {
		avail[i] = make([]int, n)
	}

	for i := 0; i < q; i++ {
		fmt.Fscan(in, &ep, &seas)
		avail[ep-1][seas-1] = 1
	}

	getMissedEpisodes(avail, out)
	out.Flush()
}

func getMissedEpisodes(avail [][]int, out *bufio.Writer) {
	for i, seas := range avail {
		for j, ep := range seas {
			if ep == 0 {
				fmt.Fprintln(out, i+1, j+1)
			}
		}
	}
}
