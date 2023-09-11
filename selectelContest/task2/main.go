package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	chestsKeys := make([][]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		keysStr := strings.ReplaceAll(scanner.Text(), "EMPTY", "-1")
		keysStrArr := strings.Split(keysStr, " ")
		keys := make([]int, len(keysStrArr))
		chestsKeys[i] = make([]int, len(keys))
		for i := range keysStrArr {
			keys[i], _ = strconv.Atoi(keysStrArr[i])
		}
		for j := 0; j < len(keys); j++ {
			chestsKeys[i][j] = keys[j]
		}
	}

	fmt.Println(isPossibleToOpenAll(0, []int{}, []int{}, chestsKeys, N))
}

func isPossibleToOpenAll(chest int, walked []int, keys []int, chestKeys [][]int, N int) bool {
	walked = append(walked, chest)

	for i := range chestKeys[chest] {
		if chestKeys[chest][i] != -1 {
			keys = append(keys, chestKeys[chest][i])
		}
	}

	for i := range keys {
		found := false
		for j := range walked {
			if walked[j] == keys[i] {
				found = true
			}
		}
		if !found {
			return isPossibleToOpenAll(keys[i], walked, keys, chestKeys, N)
		}
	}
	if len(walked) == N {
		return true
	}
	return false
}
