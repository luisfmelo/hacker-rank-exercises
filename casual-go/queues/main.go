package main

import "fmt"

func main() {
	// Simple Int Queue
	q := SimpleIntQueue{}
	fmt.Println(q.Empty())
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.Empty() == false)
	fmt.Println(q.Peek() == 1)
	fmt.Println(q.Dequeue() == 1)
	fmt.Println(q.Dequeue() == 2)
	fmt.Println(q.Dequeue() == 3)
	fmt.Println(q.Empty())

}
