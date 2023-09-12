package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	signature := [3]float64{1, 1, 1}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic("Error input")
	}

	for _, value := range Tribonacci(signature, n) {
		fmt.Print(strconv.FormatFloat(value, 'g', -1, 64) + " ")
	}
}

func Tribonacci(signature [3]float64, n int) []float64 {
	sequence := signature[:]
	for i := 3; i < n; i++ {
		sequence = append(sequence, sequence[i-1]+sequence[i-2]+sequence[i-3])
	}
	return sequence[:n]
}
