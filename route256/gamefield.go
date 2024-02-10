package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func getWidthHeight(x, y int, field *[][]rune) (int, int) {
	var width, height int

	i, j := y, x

	for i < len(*field) && (*field)[i][x] != '.' {
		height++
		i++
	}

	for j < len((*field)[y]) && (*field)[y][j] != '.' {
		width++
		j++
	}

	return width, height
}

func eraseVisitedRectangle(x, y, w, h int, field *[][]rune) {
	i, j := y, x

	for i < len(*field) && (*field)[i][x] != '.' {
		(*field)[i][x] = '.'
		(*field)[i][x+w-1] = '.'
		i++
	}
	j++
	for (*field)[y][j] != '.' {
		(*field)[y][j] = '.'
		(*field)[y+h-1][j] = '.'
		j++
	}
}

func analyzeRectangle(x, y, width, height, depth int, field *[][]rune, counts *[]int) {
	for i := y; i < y+height-2; i++ {
		for j := x; j < x+width-2; j++ {
			if (*field)[i][j] == '*' {
				//pfield := printField(*field)
				//fmt.Println(pfield)
				newWidth, newHeight := getWidthHeight(j, i, field)
				*counts = append(*counts, depth)
				analyzeRectangle(j+1, i+1, newWidth, newHeight, depth+1, field, counts)
				eraseVisitedRectangle(j, i, newWidth, newHeight, field)
			}
		}
	}
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscan(in, &n, &m)
		var field = make([][]rune, n)
		var counts []int
		for j := 0; j < n; j++ {
			var line string
			fmt.Fscan(in, &line)
			field[j] = []rune(line)
		}

		analyzeRectangle(0, 0, m, n, 0, &field, &counts)
		sort.Ints(counts)
		for j := range counts {
			fmt.Fprintf(out, "%d ", counts[j])
		}
		fmt.Fprintln(out)
	}
}
