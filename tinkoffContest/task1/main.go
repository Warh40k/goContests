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
	s, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	fmt.Println(getManualOperationsCount(s, n))

}

func getManualOperationsCount(s int, n int) int {
	for i := s; i > 0; i-- {
		if n < i {
			break
		}
		n = n - i
		if i == 1 {
			i = s + 1
		}
	}
	return n
}
