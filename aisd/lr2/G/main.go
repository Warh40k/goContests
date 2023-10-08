package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type Deque struct {
	size       int
	sum        int
	product    int
	head, tail *Node
}

func (deque *Deque) pushHead(num int) {
	temp := new(Node)
	if temp == nil {
		panic("StackOverflow")
	}
	temp.value = num

	if (*deque).size == 0 {
		(*deque).head = temp
		(*deque).tail = temp
	} else {
		(*deque).head.prev = temp
		temp.next = (*deque).head
		(*deque).head = temp
	}
	(*deque).sum += num
	(*deque).size++
}

func (deque *Deque) pushTail(num int) {
	node := new(Node)
	if node == nil {
		panic("StackOverflow")
	}
	node.value = num

	if (*deque).size == 0 {
		(*deque).head = node
		(*deque).tail = node
	} else {
		(*deque).tail.next = node
		node.prev = (*deque).tail
		(*deque).tail = node
	}
	(*deque).sum += num
	(*deque).size++
}

func (deque *Deque) popHead() int {
	if (*deque).size == 0 {
		return 0
	}
	val := (*deque).head.value
	(*deque).head = (*deque).head.next
	if deque.head != nil {
		deque.head.prev = nil
	}
	(*deque).sum -= val
	(*deque).size--

	return val
}

func (deque *Deque) popTail() int {
	if (*deque).size == 0 {
		return 0
	}
	val := (*deque).tail.value
	deque.tail = (*deque).tail.prev
	if deque.tail != nil {
		deque.tail.next = nil
	}

	(*deque).sum -= val
	(*deque).size--

	return val
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var N int
	fmt.Fscan(in, N)
	var arr = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &arr[i])
	}
	square, _ := getSquare(N, arr)
	fmt.Fprintln(out, square)
	out.Flush()
}

func getSquare(N int, arr []int) (int, error) {
	sides := make([]Deque, 4)
	var i int
	var maxmin [4]int
	for i := 0; i < N; i++ {
		sides[0].pushTail(arr[i])
	}
	square := 0
	for {
		//Вынуть из конца и вставить в начало следующего
		if i != 3 && (i == 0 || sides[i].size > sides[i+1].size) {
			val := sides[i].popTail()
			sides[i+1].pushHead(val)
		}

		// Проверка конца, обновление макс площади
		if i == 3 {
			if maxmin[0] != 0 && maxmin[1] != 0 && maxmin[2] != 0 && maxmin[3] != 0 {

				for j := 0; j < 3; j++ {
					for k := j + 1; k < 4; k++ {
						var curmax = 0

						if maxmin[j] > maxmin[k] {
							curmax = maxmin[k]
						} else {
							curmax = maxmin[j]
						}

						if maxmin[(j+2)%4] > maxmin[(k+2)%4] {
							curmax *= maxmin[(k+2)%4]
						} else {
							curmax *= maxmin[(j+2)%4]
						}

						if square < curmax {
							square = curmax
						}
					}
				}
			}
			// Если все возможные элементы просеялись вниз => все варианты просмотрены
			if sides[i].size >= N-3 {
				break
			}
			i = 0
			continue
		}
		i++
	}

	return square, nil
	//for i := 0; i < 4; i++ {
	//	result += sides[i].sum
	//}
}
