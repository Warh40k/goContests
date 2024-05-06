package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkLogin(newlogin string, oldlogins []string) bool {
	for i := range oldlogins {
		oldlogin := oldlogins[i]
		checked := false
		len1, len2 := len(newlogin), len(oldlogin)
		if len1 != len2 {
			continue
		}

		for j := 0; j < len(oldlogins); j++ {
			if oldlogin[j] != newlogin[i] {
				if !checked && oldlogin[j] == newlogin[j+1] {
					checked = true
					j++
					continue
				}
			}

		}
	}

	return true
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var nold, nnew int
	var oldlogins = make([]string, nold)
	fmt.Fscan(in, &nold)

	for i := 0; i < nold; i++ {
		fmt.Fscan(in, &oldlogins[i])
	}

	fmt.Fscan(in, &nnew)
	var newlogins = make([]string, nnew)

	for i := 0; i < nold; i++ {
		fmt.Fscan(in, &newlogins[i])
	}

}
