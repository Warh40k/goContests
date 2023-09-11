package main

import (
	"fmt"
	"time"
)

func main() {
	odd := make(chan int)
	even := make(chan int)

	go func() {
		odd <- 100
		even <- 300
	}()

	go routine(odd, true)
	go routine(even, false)

	printChan(odd)
	printChan(even)
	time.Sleep(2 * time.Second)
}

func printChan(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

func routine(buff chan int, is_odd bool) {

	k := 0
	if is_odd {
		k = 1
	}

	for i := 0; i < 10; i++ {
		if i%2 == k {
			buff <- i
		}
	}
	close(buff)
}
