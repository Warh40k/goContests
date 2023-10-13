package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	var input string
	if len(os.Args) > 1 {
		input = getInputData(os.Args[1])
	} else {
		fmt.Println("No input file specified")
		return
	}
	fmt.Println("Input text:\n", input, "\n")
	fmt.Println("Result:")
	for _, val := range findCreditCards(input) {
		fmt.Println(val)
	}
}

func findCreditCards(input string) []string {
	r, err := regexp.Compile("\\d{4}(?:\\s?\\d{4}){3}")
	if err != nil {
		panic("Regexp error")
	}
	return r.FindAllString(input, -1)
}

func getInputData(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	finfo, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bytes := make([]byte, finfo.Size())
	file.Read(bytes)
	return string(bytes)
}
