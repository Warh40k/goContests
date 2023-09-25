package main

import (
	"fmt"
)

func main() {
	var start, end uint64
	fmt.Scan(&start, &end)

	fmt.Println(getTradeDaysCount(start, end))
}

func getTradeDaysCount(start, end uint64) int {
	count := 0
	digit := 1
	// двоичное представление чисел
	bstart, bend := convertToBinary(start), convertToBinary(end)
	fmt.Println(bstart, bend)
	// количество разрядов каждого числа
	lstart, lend := len(bstart), len(bend)
	// количество старших разрядов
	constOrder := 0
	// минимальное количество групп чисел
	gCount := 1

	// Поиск старших разрядов
	if lstart == lend {
		var prev = -1
		for i := 0; bstart[i] == bend[i]; i++ {
			constOrder++
			if prev != bstart[i] {
				gCount++
			}
			prev = bstart[i]
		}
		if gCount > 3 {
			return 0
		}
	}

	digit = bend[gCount-1]

	for i := 0; i < 2; i++ {
		// Перебор
		for j := constOrder; j < lend; j++ {
			if (digit > bend[j] || digit < bstart[j]) && gCount < 3 {
				gCount++
				digit = invertDigit(digit)
				for k := j; k < lend; k++ {
					if (digit > bend[k] || digit < bstart[k]) && gCount < 3 {
						gCount++
						digit = invertDigit(digit)
						for m := k; m < lend; m++ {

						}
					} else if gCount == 3 {
						break
					} else if k == lend-1 {
						count++
					}
				}
			} else {
				break
			}
		}

		if digit == 1 {
			digit = 0
		} else {
			digit = 1
		}
	}
}

func checkRange(digit, k, lstart, lend, count, gCount int, start, end []int) int {
	if k == lend && gCount == 3 {
		count++
		return count
	} else if gCount > 3 {
		return count
	}
	if digit > end[k] || (lstart+k > lend && digit < start[]) {

	}
}

func invertDigit(digit int) int {
	if digit == 1 {
		return 0
	} else {
		return 1
	}
}

//
//func checkRange(digit, k int, start, end []int) bool {
//
//}

func convertToBinary(num uint64) []int {

	result := make([]int, 0)
	for num != 0 {
		result = append(result, int(num%2))
		num /= 2
	}

	l := len(result)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
