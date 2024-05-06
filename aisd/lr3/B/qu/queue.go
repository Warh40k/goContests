package qu

type Queue[T any] struct {
	Size     int64
	Head     *QNode[T]
	tail     *QNode[T]
	Iterator *QNode[T]
}

type QNode[T any] struct {
	Value T
	Next  *QNode[T]
}

func (q *Queue[T]) ResetIterator() {
	q.Iterator = q.Head
}

func (q *Queue[T]) Next() T {
	res := q.Iterator
	q.Iterator = q.Iterator.Next
	return res.Value
}

func (q *Queue[T]) Push(el T) {
	top := new(QNode[T])
	if top == nil {
		panic("StackOverflow")
	}

	top.Value = el
	// Если раньше элементов не было - устанавливаем head и tail на top
	if q.Size == 0 {
		q.Head = top
		q.tail = top
	} else {
		// Иначе создаем ссылку предыдущего элемента на текущий и делаем текущий эл. последним
		q.tail.Next = top
		q.tail = top
	}

	q.Size++
}

func (q *Queue[T]) Pop() T {
	if q.Size == 0 {
		panic("Underflow")
	}
	val := q.Head.Value
	q.Head = q.Head.Next
	q.Size--

	return val
}

func (priors *Queue[T]) FindPriority(k int64) T {
	var qu = priors.Head
	for i := int64(0); i < k; i++ {
		qu = qu.Next
	}
	return qu.Value
}
