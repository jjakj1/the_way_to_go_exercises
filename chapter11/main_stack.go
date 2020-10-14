package main

import (
	. "./stack"
	"fmt"
)

func main() {
	s := new(Stack)
	s.Data = make([]interface{}, 10)
	fmt.Println(s.Len())
	fmt.Println(s.IsEmpty())
	s.Push(1)
	fmt.Println(s)
	s.Pop()
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s)

	// s.Data = make([]interface{}, 10)
	// s.Push("a")
	// s.Push("b")
	// s.Pop()
	// s.Push("c")
	// fmt.Println(s)
}
