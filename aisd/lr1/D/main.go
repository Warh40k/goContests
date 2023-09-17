package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var n, m, q, seas, ep int
	fmt.Fscan(in, &n, &m, &q)

	avail := make([][]int, n)
	for i := range avail {
		avail[i] = make([]int, m)
	}

	for i := 0; i < q; i++ {
		fmt.Fscan(in, &ep, &seas)
		avail[seas-1][ep-1] = 1
	}
	missed := getMissedEpisodes(avail)

	fmt.Fprintln(out, len(missed))
	fmt.Fprintln(out, strings.Join(missed, "\n"))
	out.Flush()
}

func getMissedEpisodes(avail [][]int) []string {
	var result []string
	for i, seas := range avail {
		for j, ep := range seas {
			if ep == 0 {
				str := fmt.Sprintf("%d %d", j+1, i+1)
				result = append(result, str)
			}
		}
	}
	return result
}
