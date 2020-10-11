package main

import (
	"fmt"
)

func main() {
	str := "hello world"
	first, second := splitString(str, 5)
	fmt.Printf("first part is %s, second part is %s\n", first, second)
}

func splitString(s string, pos int) (firstPart, secondPart string) {
	b := []byte(s)
	firstPart = string(b[0:pos])
	secondPart = string(b[pos:])
	return
}
