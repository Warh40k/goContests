package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var a, b, c, d float64
	vars := [4]*float64{&a, &b, &c, &d}
	for _, val := range vars {
		sc.Scan()
		*val, _ = strconv.ParseFloat(sc.Text(), 64)
	}

	fmt.Println(solveEquation(a, b, c, d))
}

func solveEquation(a, b, c, d float64) string {
	if (a == 0 && b == 0) || (-d/c == 0) {
		return "INF"
	}
	if a == 0 && b != 0 {
		return "NO"
	}

	scope := -d / c
	x := -b / a

	if x != scope && x == float64(int(x)) {
		result := strconv.FormatInt(int64(x), 10)
		return result
	}
	return "NO"
}
