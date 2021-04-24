package main

type SimpleIntQueue []int

func (q *SimpleIntQueue) Enqueue(i int) {
	*q = append(*q, i)
}

func (q *SimpleIntQueue) Dequeue() int {
	if q.Empty() {
		panic("queue is empty")
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func (q *SimpleIntQueue) Peek() int {
	if q.Empty() {
		panic("queue is empty")
	}
	return (*q)[0]
}

func (q *SimpleIntQueue) Empty() bool {
	return len(*q) == 0
}
