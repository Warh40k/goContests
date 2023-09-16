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

	var nums []int

	for sc.Text() != "0" {
		num, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic("Error")
		}
		nums = append(nums, num)
		sc.Scan()
	}

	fmt.Println(maxNumCount(nums))
}

func maxNumCount(nums []int) int {
	var maxNum int = 0
	for _, value := range nums {
		if value > maxNum {
			maxNum = value
		}
	}
	var count = 0
	for _, value := range nums {
		if value == maxNum {
			count++
		}
	}

	return count
}
