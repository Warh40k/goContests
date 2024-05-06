package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, count int
	var result = make([]string, n/2)
	fmt.Fscan(in, &n)
	a, err := fmt.Fscan(in, nil)
	fmt.Fprintln(out, a, err)

	for i := 0; i < n; i++ {
		var num int
		fmt.Fscan(in, &num)
		if num == 1 {
			result = append(result, strconv.Itoa(count))
			count = 1
		} else {
			count++
		}
	}

	fmt.Fprintln(out, len(result))
	fmt.Fprintln(out, strings.Join(result, " "))
}
