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
	in, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic("Type mismatch. Should be integer")
	}
	fmt.Println(arabToRom(in))
}

func arabToRom(num int) string {
	seq := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	vals := map[string]int{
		"M": 1000, "CM": 900, "D": 500, "CD": 400, "C": 100, "XC": 90,
		"L": 50, "XL": 40, "X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1,
	}

	var result []string
	for i := range seq {
		val := vals[seq[i]]
		count := num / val
		num -= val * count
		for j := 0; j < count; j++ {
			result = append(result, seq[i])
		}
	}

	return strings.Join(result, "")
}
