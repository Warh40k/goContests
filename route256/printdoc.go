package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fillRange(start, end int, pages *[101]bool) {
	for i := start; i <= end; i++ {
		pages[i] = true
	}
}

func getPrintString(numpages int, input string) string {
	var printedPages [101]bool
	inputSplit := strings.Split(input, ",")
	for i := range inputSplit {
		pageRange := strings.Split(inputSplit[i], "-")
		if len(pageRange) == 2 {
			start, _ := strconv.Atoi(pageRange[0])
			end, _ := strconv.Atoi(pageRange[1])
			fillRange(start, end, &printedPages)
		} else {
			page, _ := strconv.Atoi(pageRange[0])
			printedPages[page] = true
		}
	}
	var result []string
	for i := 1; i <= numpages; i++ {
		count := 0
		var first int
		last := 0
		var element string

		for j := i; j <= numpages; j++ {
			if printedPages[j] == false {
				count++
				if first == 0 {
					first = j
				}
				last = j
			} else if first != 0 {
				break
			}
		}
		if count == 0 {
			break
		} else if count > 1 {
			element = fmt.Sprintf("%d-%d", first, last)
		} else if count == 1 {
			element = strconv.Itoa(first)
		}
		i = last
		result = append(result, element)
	}
	return strings.Join(result, ",")
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var l int
	fmt.Fscan(in, &l)

	for i := 0; i < l; i++ {
		var n int
		var s string
		fmt.Fscan(in, &n, &s)
		fmt.Fprintln(out, getPrintString(n, s))
	}
}
