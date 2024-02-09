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

func countPlayerValue(player [2]string, set1, set2 []string) (int, []string) {
	val1, val2 := cardValues[string(player[0][0])], cardValues[string(player[1][0])]

	ratio1, ratio2 := 1, 1
	if player[0][0] == player[1][0] {
		ratio1++
		ratio2++
	}
	if len(set1) != 0 {
		ratio1++
	}
	if len(set2) != 0 {
		ratio2++
	}

	setVal1, setVal2 := val1*ratio1, val2*ratio2
	maxVal := setVal1
	maxSet := set1
	if setVal2 > setVal1 {
		maxVal = setVal2
		maxSet = set2
	}

	return maxVal, maxSet
}

func getWinCards(playersCards [][2]string) {
	fp := playersCards[0]
	fpSet1, fpSet2 := getMaxSet(string(fp[0][1])), getMaxSet(string(fp[1][1]))

	fpValue, fpSet := countPlayerValue(fp, fpSet1, fpSet2)

	for i := 1; i < len(playersCards); i++ {
		ratio1, ratio2 := 1, 1
		player := playersCards[i]
		maxScore := 0
		for j := 0; j < len(fpSet); j++ {
			if player[0][0] == fpSet[0][0] {

			}
		}
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
