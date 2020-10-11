package main

import (
	"fmt"
)

func main() {
	s := "hello"
	b := []byte(s)
	b[0] = 'c'
	s2 := string(b)
	fmt.Println(s)
	fmt.Println(s2)
}
