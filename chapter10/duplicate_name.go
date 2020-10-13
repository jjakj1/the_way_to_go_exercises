package main

import (
	"fmt"
)

type A struct{ a int }

type B struct{ a, b int }

type C struct {
	A
	B
}

func main() {
	c := new(C)
	fmt.Println(c.A.a)
	fmt.Println(c.B.a)
}
