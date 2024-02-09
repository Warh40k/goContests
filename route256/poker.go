package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	cardCount = 52
)

var cardValues = map[string]int{"2": 0, "3": 1, "4": 2, "5": 3, "6": 4, "7": 5, "8": 6, "9": 7, "T": 8, "J": 9, "Q": 10, "K": 11, "A": 12}
var cardMast = []string{"S", "C", "D", "H"}

func getMaxSet(cardVal string, cardsOnHands map[string]bool) []string {
	var result []string

	for j := range cardMast {
		card := cardVal + cardMast[j]
		if _, ok := cardsOnHands[card]; ok == true {
			result = append(result, card)
		}
	}

	return result
}

func countPlayerValue(player [2]string, sets [2][]string) [2]int {
	var vals = [2]int{cardValues[string(player[0][0])], cardValues[string(player[1][0])]}

	var ratios = [2]int{}
	var setVals = [2]int{}

	if player[0][0] == player[1][0] {
		ratios[0]++
		ratios[1]++
	}

	for i := range sets {
		if len(sets[i]) != 0 {
			ratios[i]++
		}
		setVals[i] = vals[i] * ratios[i]
	}

	return setVals
}

func getWinCards(playersCards [][2]string, cardsOnHands map[string]bool) []string {
	fp := playersCards[0]
	var fpSets = [2][]string{getMaxSet(string(fp[0][1]), cardsOnHands), getMaxSet(string(fp[1][1]), cardsOnHands)}

	fpValues := countPlayerValue(fp, fpSets)

	for i := 1; i < len(playersCards); i++ {
		ratio1, ratio2 := 1, 1
		player := playersCards[i]
		for j := 0; j < len(fpSets); j++ {
			if ratio1 != 1 && player[0][0] == fpSets[j][0] {
				ratio1++
			}
			if ratio2 != 1 && player[1][0] == fpSets[j][0] {
				ratio2++
			}
		}
		if player[0][0] == player[1][0] {
			ratio1++
			ratio2++
		}

		val1, val2 := cardValues[string(player[0][0])], cardValues[string(player[1][0])]
		setVal1, setVal2 := val1*ratio1, val2*ratio2
	}
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var l int
	fmt.Fscan(in, &l)

	for i := 0; i < l; i++ {
		var n int
		playersCards := make([][2]string, l)
		var cardsOnHands = map[string]bool{}
		fmt.Fscan(in, &n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &playersCards[j][0], &playersCards[j][1])
			cardsOnHands[playersCards[j][0]] = true
			cardsOnHands[playersCards[j][1]] = true
		}
		getWinCards(playersCards)
	}
}
