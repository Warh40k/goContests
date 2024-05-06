package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n, t int
	var alph = make(map[string]int)
	fmt.Fscan(in, &n, &t)

	for i := 0; i < n; i++ {
		var sym string
		fmt.Fscan(in, &sym)
		alph[sym]++
	}

	for i := 0; i < t; i++ {
		var S string
		fmt.Fscan(in, &S)
		if checkPass(S, n, alph) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func checkPass(S string, n int, alph map[string]int) bool {
	var cur = make(map[string]int)
	for _, sym := range S {
		_, ok := alph[string(sym)]
		_, ok2 := cur[string(sym)]
		if (ok2 && cur[string(sym)] >= alph[string(sym)]) || !ok {
			return false
		}
		cur[string(sym)]++
		n--
	}
	if n != 0 {
		return false
	}
	return true
}
