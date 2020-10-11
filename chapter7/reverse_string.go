package main

import (
	"fmt"
)

func main() {
	s := "Google"
	fmt.Println(reverseString(s))
}

func reverseString(str string) (res string) {
	b := []byte(str)
	l := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[l-i-1] = b[l-i-1], b[i]
	}
	res = string(b)
	return
}
