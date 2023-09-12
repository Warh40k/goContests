package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic("Input error")
	}
	fmt.Println(FindNb(n))
}

func FindNb(m int) int {
	for i := 0; m > 0; i++ {
		m = m - i*i*i
		if m == 0 {
			return i
		}
	}
	return -1
}
