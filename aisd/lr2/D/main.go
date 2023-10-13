package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	size int
	head *QNode
	tail *QNode
}

type QNode struct {
	value byte
	next  *QNode
}

func (q *Queue) enqueue(num byte) {
	top := new(QNode)
	if top == nil {
		panic("StackOverflow")
	}

	top.value = num
	// Если раньше элементов не было - устанавливаем head и tail на top
	if q.size == 0 {
		q.head = top
		q.tail = top
	} else {
		// Иначе создаем ссылку предыдущего элемента на текущий и делаем текущий эл. последним
		q.tail.next = top
		q.tail = top
	}

	q.size++
}

func (q *Queue) dequeue() byte {
	if q.size == 0 {
		panic("Underflow")
	}
	val := q.head.value
	q.head = q.head.next
	q.size--

	return val
}

type Stack struct {
	size, product int
	top           *SNode
}

type SNode struct {
	next  *SNode
	value int
}

func (s *Stack) push(num int) {
	temp := new(SNode)
	if temp == nil {
		panic("StackOverflow")
	}
	temp.value = num
	temp.next = s.top
	s.top = temp

	s.product *= num
	s.size++
}

func (s *Stack) pop() int {
	if s.size == 0 {
		panic("Underflow")
	}
	val := s.top.value
	s.top = s.top.next

	s.product /= val
	s.size--

	return val
}

func main() {
	var input string
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	fmt.Fscan(in, &input)

	fmt.Fprintln(out, parseFormula(input))
	out.Flush()
}

func parseFormula(formula string) string {
	// Количества элементов в формуле (от A до Z)
	var elems [26]int
	var elemCount int
	length := len(formula)
	factors := new(Stack)
	factors.product = 1

	for i := 0; i < length; i++ {
		// Если это название элемента => добавляем в стек дефолтное количество
		if formula[i] >= 65 && formula[i] <= 90 {
			elems[formula[i]-65] += factors.product
			elemCount++
		} else if formula[i] == 40 {
			// Если это открывающая скобка => ищем закрывающую
			j := findCloseBracket(formula, i, length)
			// Ищем коэффициент после скобки
			snum, _ := getNum(formula, j, length)
			// Добавляем коэффициент в стек
			factor, _ := strconv.Atoi(snum)
			factors.push(factor)
		} else if formula[i] == 41 {
			// Если это закрывающая скобка
			// => удаляем множитель и пропускаем его
			factor := factors.pop()
			i += len(strconv.Itoa(factor))
		} else {
			// Иначе добавляем количество к предудущему элементу
			snum, numSize := getNum(formula, i, length)
			num, _ := strconv.Atoi(snum)
			elems[formula[i-1]-65] += (num - 1) * factors.product
			i += numSize - 1
		}
	}

	// Собираем отсортированную формулу
	result := make([]string, elemCount)

	for i, j := 0, 0; i < 26; i++ {
		if elems[i] != 0 {
			count := ""
			if elems[i] != 1 {
				count = strconv.Itoa(elems[i])
			}
			result[j] = string(rune(i+65)) + count
			j++
		}
	}
	return strings.Join(result, "")
}

func findCloseBracket(formula string, i int, length int) int {
	balance := 1
	j := i + 1
	for balance != 0 {
		if formula[j] == 40 {
			// Если это открывающая скобка
			balance++
		} else if formula[j] == 41 {
			// Закрывающая скобка
			balance--
		}
		// Если последний элемент
		if j == length-1 {
			break
		}
		j++
	}
	return j
}

func getNum(formula string, k, length int) (string, int) {
	nums := new(Queue)
	for i := k; i < length; i++ {
		if formula[i] >= 48 && formula[i] <= 57 {
			nums.enqueue(formula[i])
		} else if nums.size == 0 {
			nums.enqueue(49)
			break
		} else {
			break
		}
	}
	result := make([]byte, nums.size)
	elem := nums.head
	for i := range result {
		result[i] = elem.value
		elem = elem.next
	}
	return string(result), nums.size
}
