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
	//testHard(out)
	//os.Exit(0)
	testCorrect()
	os.Exit(0)
	var N int
	fmt.Fscan(in, &N)
	var X = make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(in, &X[i])
	}
	count := countBorders(N, X)
	//countOld, _ := countBordersOld(N, X)
	fmt.Fprintln(out, count)
	out.Flush()
}

func countBorders(N int, X []int) int {
	count := 0
	marked := make([][]int, 2)
	marked[0] = make([]int, N)
	marked[1] = make([]int, N)

	table := make([][]int, N)
	for i := 0; i < N; i++ {
		table[i] = make([]int, N)
		for j := 0; j < N; j++ {
			table[i][j] = X[i] & X[j]
		}
	}
	for k := 1; k < N; k++ {
		counter := 0

		for i, j := 0, k; i < N; i, j = i+1, j+1 {
			if j == N {
				j = 0
			}

			if X[i]&X[j] != 0 {
				counter++
				if k == 1 {
					count++
				} else if marked[0][i] == 0 && X[(i+N+1)%N]&X[j] == 0 {
					count++
					marked[0][(i+N+1)%N] = 1
				}

				marked[0][i] = 1
			} else if k != 1 && (marked[0][i] != 0 || marked[0][(i+N+1)%N] != 0) {
				marked[0][i] = 1
				counter++
			}
		}
		if counter >= N {
			break
		}
		//marked[0], marked[1] = marked[1], marked[0]
	}

	if count == 0 {
		return 1
	}
	return count
}

func testHard(writer *bufio.Writer) {
	N := int(4 * math.Pow(10, 5))
	arr := make([]int, N)

	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(int(math.Pow(10, 9)) + 1)
	}
	count := countBorders(N, arr)
	fmt.Fprintln(writer, count)
	writer.Flush()
}

func testCorrect() {
	for {
		N := rand.Intn(5) + 1
		arr := make([]int, N)

		for i := 0; i < N; i++ {
			arr[i] = rand.Intn(10) + 1
		}
		count1, _ := countBordersOld(N, arr)
		count2 := countBorders(N, arr)

		if count1 != count2 {
			fmt.Println(count1, count2, arr)
			return
		}
		//fmt.Println("ok")
	}
}

func countBordersOld(N int, X []int) (int, []int) {
	count := 0
	var borders []int
	for i := 1; i < N; i++ {
		for j := 0; j < N; j++ {
			var next = j + i
			if next > N-1 {
				next -= N
			}
			if (i == 1 || hasBordersBetween(j, i, N, borders)) && X[j]&X[next] != 0 {
				borders = append(borders, next)
				count++
			}
		}
	}
	if count == 0 {
		count = 1
	}
	return count, borders
}
func hasBordersBetween(start, step, N int, borders []int) bool {
	if borders != nil {
		for _, val := range borders {
			if (start+step <= N-1 && val > start && val <= start+step) ||
				(start+step > N-1 && (val > start || val <= start+step-N)) {
				return false
			}
		}
	}
	return true
}
