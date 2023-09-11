package main

import "fmt"

func main() {
	a := incFunc()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}

func incFunc() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
