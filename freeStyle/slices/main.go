package main

import "fmt"

func main() {
	a := make([]string, 3)
	a[0] = "3"
	a[1] = "4"

	a = append(a, "hello")
	fmt.Println(a)
}
