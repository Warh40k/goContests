package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		go ahello(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("this is the end")

}

func ahello(i int) {
	fmt.Println(i)
}
