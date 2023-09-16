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
	sc.Scan()
	troom, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic("Err")
	}
	sc.Scan()
	tcond, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic("Err")
	}
	sc.Scan()
	mode := sc.Text()
	fmt.Println(setTemp(troom, tcond, mode))
}

func setTemp(troom int, tcond int, mode string) int {
	if mode == "auto" ||
		(mode == "freeze" && troom > tcond) ||
		(mode == "heat" && troom < tcond) {
		return tcond
	}
	return troom
}
