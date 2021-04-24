package main

import "fmt"

func main() {
	// Simple Int Stack
	s := SimpleIntStack{}
	fmt.Println(s.Empty())
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Empty() == false)
	fmt.Println(s.Peek() == 3)
	fmt.Println(s.Pop() == 3)
	fmt.Println(s.Pop() == 2)
	fmt.Println(s.Pop() == 1)
	fmt.Println(s.Empty())

}
