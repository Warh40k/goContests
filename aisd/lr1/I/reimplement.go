package main

import (
	"fmt"
	"math"
)

func main() {
	var start, end uint64
	fmt.Scan(&start, &end)

	fmt.Println(getTradeDaysCount(start, end), countTradeDays(start, end))
}

func getTradeDaysCount(start, end uint64) int {
	count := 0
	// двоичное представление чисел
	bstart, bend := convertToBinary(start), convertToBinary(end)
	// количество разрядов каждого числа
	lstart, lend := len(bstart), len(bend)

	if lstart != lend {
		bstart = resizeNum(lstart, lend, bstart)
	}

	var prev, gcount int = -1, 0
	//Подсчет изначального количества групп
	for i := range bstart {
		if bstart[i] != prev {
			prev = bstart[i]
			gcount++
		}
	}
	if gcount == 3 {
		count++
	}

	iStart, jStart := start, start

	//for i := lend - 1; i > 0; i``-- {
	//	if bstart[i] == 0 {
	//		iStart += uint64(math.Pow(2, float64(i)))
	//		if iStart > end {
	//			break
	//		}
	//		count++
	//	} else {
	//		count++
	//	}
	//	jStart = start
	//	for j := 0; j < i-1; j++ {
	//		if bstart[j] == 0 {
	//			jStart += uint64(math.Pow(2, float64(i)))
	//			if jStart > end {
	//				break
	//			}
	//			count++
	//		} else {
	//			count++
	//		}
	//	}
	//
	//}
	for i := lstart - 1; i < lend; i++ {
		for j := 0; j < i-1; j++ {
			if bstart[j] == 0 {
				jStart += uint64(math.Pow(2, float64(j)))
				if jStart > end {
					break
				}
				count++
			}
		}
	}
	fmt.Println(bstart, bend)
	return count
}

func convertToBinary(num uint64) []int {

	result := make([]int, 0)
	for num != 0 {
		result = append(result, int(num%2))
		num /= 2
	}
	//
	//l := len(result)
	//for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
	//	result[i], result[j] = result[j], result[i]
	//}
	return result
}

func resizeNum(len1, len2 int, num []int) []int {
	result := make([]int, len2)
	for i := 0; i < len1; i++ {
		result[i] = num[i]
	}

	return result
}

func countTradeDays(start, end uint64) int {
	count := 0
	for i := start; i <= end; i++ {
		if checkBinary(i) {
			count++
		}
	}
	return count
}

func checkBinary(num uint64) bool {
	prev := num % 2
	var cur uint64 = 0
	factor := num / 2
	gcount := 0

	for factor != 0 {
		cur = factor % 2
		if cur != prev {
			gcount++
			prev = cur
		}
		if gcount > 2 {
			return false
		}
		if factor != 0 {
			factor = factor / 2
		}
	}
	if gcount != 2 {
		return false
	}
	return true
}
