package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
)

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	test(out)
	os.Exit(0)
	var N int
	fmt.Fscan(in, &N)
	var X = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &X[i])
	}
	count := countBorders(N, X)
	fmt.Fprintln(out, count)
	out.Flush()
}

func countBorders(N int, X []int) int {
	count := 0
	marked := make([][]int, N)
	//table := make([][]int, N)
	for i := 0; i < N; i++ {
		marked[i] = make([]int, N)
		//table[i] = make([]int, N)
		for j, k := i+1, 0; k < N; j++ {
			if j == N {
				j = 0
			}
			//table[i][k] = X[i] & X[j]
			k++
		}
	}

	for k := 0; k < N; k++ {

		for i := 0; i < N; i++ {
			prev1 := &marked[i][int(math.Abs(float64((j-1)%N)))]
			prev2 := &marked[int(math.Abs(float64((i+1)%N)))][int(math.Abs(float64((j-1)%N)))]
			if *prev1 == 0 && *prev2 == 0 {
				if table[i][j] != 0 {
					count++
					marked[i][j], *prev1, *prev2 = 1, 1, 1
				}
			} else {
				marked[i][j] = 1
			}
		}
	}

	if count == 0 {
		return 1
	}
	return count
}
func hasBorders(start, end, N int, table [][]int) bool {
	for k := end - 1; k != start; k-- {
		if k == -1 {
			k = N - 1
		}
		i := start
		j := k
		for j != end+1 {
			if i == N {
				i = 0
			}
			if j == N {
				j = 0
			}
			if table[i][j] != 0 {
				return true
			}
			i++
			j++
		}
	}
	return false
	//for i != end {
	//	j := end - 1
	//
	//	for j != i {
	//		if j == -1 {
	//			j = N - 1
	//		}
	//		if table[i][j] != 0 {
	//			return j
	//		}
	//		j--
	//	}
	//	if i == N-1 {
	//		i = 0
	//	} else {
	//		i++
	//	}
	//}
}

func test(writer *bufio.Writer) {
	N := int(4 * math.Pow(10, 5))
	arr := make([]int, N)

	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(int(math.Pow(10, 9)) + 1)
	}
	count := countBorders(N, arr)
	fmt.Fprintln(writer, count)
	writer.Flush()
}
