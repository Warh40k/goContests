package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var exprs []string

	for scanner.Scan() {
		exprs = append(exprs, scanner.Text())
	}

	fmt.Println(exprs)

	for _, expr := range exprs {
		go evaluate(expr)
	}

	time.Sleep(2 * time.Second)
}

func evaluate(expr string) {
	re, _ := regexp.Compile("[+\\-*/]")
	matched := re.FindString(expr)
	elems := strings.Split(expr, matched)

	if len(elems) < 2 {
		return
	}

	num1, _ := strconv.Atoi(elems[0])
	num2, _ := strconv.Atoi(elems[1])

	result := expr + "="

	switch matched {
	case "*":
		result += strconv.Itoa(num1 * num2)
	case "+":
		result += strconv.Itoa(num1 + num2)
	case "-":
		result += strconv.Itoa(num1 + num2)
	case "/":
		result += strconv.Itoa(num1 / num2)
	}

	fmt.Println(result)
}
