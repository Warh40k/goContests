package main

import "fmt"

func main() {
	hanoiTowers(4, 1, 3)
}

func hanoiTowers(n int, fromPeg int, toPeg int) {
	if n == 1 {
		fmt.Printf("Move from peg %d to peg %d\n", fromPeg, toPeg)
		return
	}
	unusedPeg := 6 - fromPeg - toPeg
	hanoiTowers(n-1, fromPeg, unusedPeg)
	fmt.Printf("Move from peg %d to peg %d\n", fromPeg, toPeg) // Двигает нижний диск в конечное место
	hanoiTowers(n-1, unusedPeg, toPeg)
}
