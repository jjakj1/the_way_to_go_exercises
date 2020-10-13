package main

import (
	"./stack"
	"fmt"
)

func main() {
	st := stack.New(4)
	st.Push(1)
	st.Push(2)
	fmt.Println(st)
	st.Pop()
	fmt.Println(st)
	st.Push(3)
	st.Push(4)
	fmt.Println(st)
}
