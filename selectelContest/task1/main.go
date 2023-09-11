package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Number struct {
	count int
	value int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	params := strings.Split(scanner.Text(), " ")
	N, _ := strconv.Atoi(params[0])
	k, _ := strconv.Atoi(params[1])

	scanner.Scan()
	num_string := strings.Split(scanner.Text(), " ")
	nums := make([]int, N)
	for i := range num_string {
		nums[i], _ = strconv.Atoi(num_string[i])
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	getTopEntries(k, nums, N)
}

func getTopEntries(k int, nums []int, N int) {
	var counts []Number
	for i := 0; i < N; i++ {
		num := Number{0, nums[i]}
		for j := i; j < N; j++ {
			if nums[j] == nums[i] {
				num.count++
			} else {
				counts = append(counts, num)
				i = j - 1
				break
			}
		}
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i].count > counts[j].count
	})

	counts = counts[:k]

	sort.Slice(counts, func(i, j int) bool {
		return counts[i].value < counts[j].value
	})

	var result []string
	for i := range counts {
		result = append(result, strconv.Itoa(counts[i].value))
	}
	fmt.Println(strings.Join(result, " "))
}
