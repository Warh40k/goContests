package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(b)

	fmt.Println(getValueArithmetics(b))
}

func getValueArithmetics(a *int) *int {
	*a = *a + 3
	return a
}
