package main

import (
	"fmt"
)

type T struct {
	a int
	b float32
	c string
}

func main() {
	t := &T{7, -2.35, "abcd\\tdef"}
	fmt.Println(t)
}

func (t *T) String() string {
	return fmt.Sprintf("%d / %f / %s", t.a, t.b, t.c)
}
