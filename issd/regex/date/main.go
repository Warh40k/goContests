package main

import (
	"fmt"
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
	"os"
)

func main() {
	var input string
	if len(os.Args) > 1 {
		input = getInputData(os.Args[1])
	} else {
		fmt.Println("No input file specified")
		return
	}

	fmt.Println("Input text:\n", input)
	fmt.Println("\nResult:")
	fmt.Println(findDates(input))
}

func findDates(input string) string {
	r, _ := pcre.Compile(`^(?:(?:(?:0?[13578]|1[02])(\/|-|\.)31)\1|(?:(?:0?[1,3-9]|1[0-2])(\/|-|\.)(?:29|30)\2))(?:(?:1[6-9]|[2-9]\d)?\d{2})$|^(?:0?2(\/|-|\.)29\3(?:(?:(?:1[6-9]|[2-9]\d)?(?:0[48]|[2468][048]|[13579][26])|(?:(?:16|[2468][048]|[3579][26])00))))$|^(?:(?:0?[1-9])|(?:1[0-2]))(\/|-|\.)(?:0?[1-9]|1\d|2[0-8])\4(?:(?:1[6-9]|[2-9]\d)?\d{2})$`, 0)
	m := r.MatcherString(input, -1)
	return m.GroupString(0)
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
