package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")
	k, _ := strconv.Atoi(inputs[1])
	scanner.Scan()
	p := scanner.Text()

	counts := countEntries(p)
	fmt.Println(countUniqs(k, counts))
}

func countEntries(p string) []int {
	countMap := make(map[string]int)
	var counts []int

	for i := range p {
		sym := string(p[i])
		countMap[sym] = 0
		for j := range p {
			if p[i] == p[j] {
				countMap[sym]++
			}
		}
	}

	for _, v := range countMap {
		counts = append(counts, v)
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i] < counts[j]
	})
	return counts
}

func countUniqs(k int, counts []int) int {
	remCount := 0
	for i := range counts {
		if k-counts[i] >= 0 {
			k = k - counts[i]
			remCount++
		} else {
			break
		}
	}
	return len(counts) - remCount
}
