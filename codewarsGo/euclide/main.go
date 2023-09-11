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
	x, err := strconv.ParseUint(scanner.Text(), 10, 32)
	if err != nil {
		fmt.Println("Ошибка ввода")
	}

	scanner.Scan()
	y, err := strconv.ParseUint(scanner.Text(), 10, 32)
	if err != nil {
		fmt.Println("Ошибка ввода")
	}

	xi := uint32(x)
	yi := uint32(y)

	fmt.Println(Gcd(xi, yi))
}

func Gcd(x, y uint32) uint32 {
	if x-y == 0 {
		return x
	}

	var nmin uint32
	var nmax uint32
	if x > y {
		nmax = x
		nmin = y
	} else {
		nmax = y
		nmin = x
	}

	if nmax-nmin < nmax {
		return Gcd(nmin, nmax-nmin)
	}

	return Gcd(nmax-nmin, nmin)
}
