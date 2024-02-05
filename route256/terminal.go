package main

import (
	"bufio"
	"fmt"
	"os"
)

const inputStart = 128512

type Queue struct {
	size, linePos, curLine, lines int
	cur, head, tail               *QNode
}

type QNode struct {
	value rune
	next  *QNode
	prev  *QNode
}

func NewQueue() *Queue {
	q := &Queue{
		curLine: 1,
		lines:   1,
	}
	qnode := &QNode{
		value: inputStart,
		next:  nil,
		prev:  nil,
	}
	q.head = qnode
	q.tail = qnode
	q.cur = qnode

	return q
}

func (q *Queue) moveLeft() {
	if q.cur.value != inputStart && q.cur.value != '\n' {
		q.cur = q.cur.prev
		q.linePos--
	}
}

func (q *Queue) moveRight() {
	if q.cur.next != nil && q.cur.next.value != '\n' {
		q.cur = q.cur.next
		q.linePos++
	}
}

func (q *Queue) moveUp() {
	if q.curLine-1 <= 0 {
		return
	}
	oldPos := q.linePos
	q.goHome()
	q.cur = q.cur.prev
	q.goHome()
	newPos := 0
	for i := 0; i < oldPos; i++ {
		if q.cur.next == nil || q.cur.next.value == '\n' {
			break
		} else {
			q.cur = q.cur.next
		}
		newPos++
	}
	q.linePos = newPos
	q.curLine--
}

func (q *Queue) moveDown() {
	if q.curLine+1 > q.lines {
		return
	}
	oldPos := q.linePos
	q.goEnd()
	q.cur = q.cur.next
	newPos := 0
	for i := 0; i < oldPos; i++ {
		if q.cur.next == nil || q.cur.next.value == '\n' {
			break
		}
		newPos++
		q.cur = q.cur.next
	}
	q.linePos = newPos
	q.curLine++
}

func (q *Queue) goHome() {
	for q.cur.value != inputStart && q.cur.value != '\n' {
		q.cur = q.cur.prev
	}
	q.linePos = 0
}

func (q *Queue) goEnd() {
	for q.cur.next != nil && q.cur.next.value != '\n' {
		q.cur = q.cur.next
		q.linePos++
	}
}

func (q *Queue) newLine() {
	q.insert('\n')
	q.cur = q.cur.next
	q.linePos = 0
	q.curLine++
	q.lines++
}

func (q *Queue) insert(char rune) {
	node := new(QNode)
	node.value = char
	node.next = q.cur.next
	node.prev = q.cur
	if node.next != nil {
		node.next.prev = node
	}
	q.cur.next = node
	q.moveRight()

	q.size++
}

func getOutput(input string) string {
	var term = NewQueue()
	var out = make([]rune, 0, 50)
	for i := range input {
		var char = rune(input[i])
		switch char {
		case 'L':
			term.moveLeft()
		case 'R':
			term.moveRight()
		case 'B':
			term.goHome()
		case 'E':
			term.goEnd()
		case 'N':
			term.newLine()
		case 'U':
			term.moveUp()
		case 'D':
			term.moveDown()
		default:
			term.insert(char)
		}
	}
	pointer := term.head.next
	for i := 0; i < term.size; i++ {
		out = append(out, pointer.value)
		pointer = pointer.next
	}
	return string(out)
}

func main() {
	in, out := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var input string
		fmt.Fscan(in, &input)
		fmt.Fprintln(out, getOutput(input))
		fmt.Fprintln(out, "-")
	}
	out.Flush()
}
