package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	cardCount = 52
)

var cardValues = map[string]int{"2": 1, "3": 2, "4": 3, "5": 4, "6": 5, "7": 6, "8": 7, "9": 8, "T": 9, "J": 10, "Q": 11, "K": 12, "A": 13}
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

func countPlayerValue(player [2]string, card string) int {
	var handVals = [2]int{cardValues[string(player[0][0])], cardValues[string(player[1][0])]}
	var cardVal = cardValues[string(card[0])]

	var combos = [2]int{}
	var setVals = [2]int{}

	if player[0][0] == player[1][0] {
		combos[0]++
		combos[1]++
	}

	for i := range player {
		if player[i][0] == card[0] {
			combos[i]++
		}
		if combos[i] == 0 {
			setVals[i] = handVals[i]
			if cardVal > handVals[i] {
				setVals[i] = cardVal
			}
		} else {
			setVals[i] = handVals[i] + 14*combos[i]
		}
	}
	maxVal := setVals[0]
	if setVals[1] > setVals[0] {
		maxVal = setVals[1]
	}
	return maxVal
}

func checkWin(playersCards [][2]string, card string, fpVal int) bool {
	for j := 1; j < len(playersCards); j++ {
		player := playersCards[j]
		playerVal := countPlayerValue(player, card)
		if playerVal > fpVal {
			return false
		}
	}
	return true
}

func getWinCards(playersCards [][2]string, cardsOnHands map[string]bool) []string {
	fp := playersCards[0]
	var winCards []string
	for val, _ := range cardValues {
		for i := range cardMast {
			card := val + cardMast[i]
			if _, ok := cardsOnHands[card]; ok == true {
				continue
			}
			fpVal := countPlayerValue(fp, card)
			if checkWin(playersCards, card, fpVal) {
				winCards = append(winCards, card)
			}
		}
	}
	return winCards
}

func main() {
	//file, err := os.Create("result.poker")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()

	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var l int
	fmt.Fscan(in, &l)

	for i := 0; i < l; i++ {
		var n int
		var cardsOnHands = map[string]bool{}
		fmt.Fscan(in, &n)
		playersCards := make([][2]string, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &playersCards[j][0], &playersCards[j][1])
			cardsOnHands[playersCards[j][0]] = true
			cardsOnHands[playersCards[j][1]] = true
		}
		winCards := getWinCards(playersCards, cardsOnHands)
		fmt.Fprintln(out, len(winCards))
		for j := range winCards {
			fmt.Fprintln(out, winCards[j])
		}
	}
}
