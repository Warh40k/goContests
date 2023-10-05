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
	marked := make([]int, N)
	for k := 1; k < N; k++ {
		counter, success := 0, false

		for i, j := 0, k; i < N; i, j = i+1, j+1 {
			if j == N {
				j = 0
			}

			if j == 1e5 {
				fmt.Print("Capez")
			}
			if X[i]&X[j] != 0 {
				counter++
				if k == 1 || (marked[i] == 0 && marked[(i+N+1)%N] == 0) && success == false {
					count++
					success = true
				} else {
					success = false
				}
				marked[i] = 1
			} else if k != 1 && (marked[i] != 0 || marked[(i+N+1)%N] != 0) {
				marked[i] = 1
				counter++
				success = false
			} else {
				marked[i] = 0
			}
		}
		if counter >= N {
			break
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
