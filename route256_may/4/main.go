package main

import (
	"bufio"
	"fmt"
	"os"
)

var maxValue float64

func swapDollarCurrencies(i int, banks [][][]float64) {
	temp := banks[i][2]
	banks[i][2] = banks[i][3]
	banks[i][3] = temp
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int

	fmt.Fscan(in, &t)

	for m := 0; m < t; m++ {
		var banks = make([][][]float64, 3)
		var visited = make([]bool, 3)
		maxValue = 0
		for i := 0; i < 3; i++ {
			banks[i] = make([][]float64, 6)
			for j := 0; j < 6; j++ {
				banks[i][j] = make([]float64, 2)
				fmt.Fscan(in, &banks[i][j][0], &banks[i][j][1])
			}
			swapDollarCurrencies(i, banks)
		}

		for i := 0; i < 3; i++ {
			searchBanks(i, [3]float64{1, 0, 0}, visited, banks)
		}

		fmt.Fprintln(out, maxValue)
	}
}

func searchBanks(curBank int, balance [3]float64, visited []bool, banks [][][]float64) {
	if balance[1] > maxValue {
		maxValue = balance[1]
	}
	for i := range banks {
		if visited[i] {
			continue
		}
		for j := range balance {
			if balance[j] == 0 {
				continue
			}
			for k := 0; k < 2; k++ {
				temp := balance[j]
				balance[(j+1+k)%3] = balance[j] / banks[i][2*j+k][0] * banks[i][2*j+k][1]
				balance[j] = 0
				visited[i] = true
				searchBanks(i, balance, visited, banks)
				balance[j] = temp
				balance[(j+1+k)%3] = 0
			}
		}
	}
	visited[curBank] = false
}
