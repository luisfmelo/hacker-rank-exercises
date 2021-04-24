package main

type SimpleIntStack []int

func (q *SimpleIntStack) Push(v int) {
	*q = append(*q, v)
}

func (q *SimpleIntStack) Pop() int {
	if q.Empty() {
		panic("stack is empty")
	}
	v := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return v
}

func (q *SimpleIntStack) Peek() int {
	if q.Empty() {
		panic("stack is empty")
	}
	return (*q)[len(*q)-1]
}

func (q *SimpleIntStack) Empty() bool {
	return len(*q) == 0
}
