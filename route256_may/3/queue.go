package main

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

type Node struct {
	value rune
	next  *Node
	prev  *Node
}

func (l *LinkedList) PushFront(value rune) {
	node := &Node{value: value}
	if l.head == nil {
		l.head, l.tail = node, node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	l.len++
}

func (l *LinkedList) PushBack(value rune) {
	node := &Node{value: value}
	if l.tail == nil {
		l.head, l.tail = node, node
	} else {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	}
	l.len++
}

func (l *LinkedList) InsertAfter(node *Node, value rune) {
	newNode := &Node{value: value}
	if node == l.tail {
		l.tail = newNode
	} else {
		newNode.next, newNode.prev = node.next, node
		node.next.prev, node.next = newNode, newNode
	}
	l.len++
}

func NewLinkedList(s rune, len int) *LinkedList {
	list := &LinkedList{}
	for i := 0; i < len; i++ {
		list.PushBack(s)
	}
	return list
}
