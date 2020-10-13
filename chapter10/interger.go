package main

import "fmt"

// type Integer int
type Integer struct {
	n int
}

func main() {
	i := &Integer{1}
	fmt.Println(i.get())
	f(i.get())
}

// func (i Integer) get() int {
// 	return int(i)
// }

func (i Integer) get() int {
	return i.n
}

func f(i int) {
	fmt.Println(i + 1)
}
