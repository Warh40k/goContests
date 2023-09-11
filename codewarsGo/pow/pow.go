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
	a, _ := strconv.ParseFloat(scanner.Text(), 32)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	power(a, n, 1)
}

func power(a float64, n int, result float64) {
	if n == 0 {
		fmt.Println(result)
		return
	}
	result = result * a
	power(a, n-1, result)
}
