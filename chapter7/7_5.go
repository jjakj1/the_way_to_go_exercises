package main

import (
	"fmt"
)

func main() {
	slice := []byte{'a', 'b', 'c'}
	data := []byte{'d', 'e'}
	fmt.Println(Append(slice, data))
}

func Append(slice, data []byte) []byte {
	s := make([]byte, len(slice)+len(data))
	copy(s, slice)
	copy(s[len(slice):], data)
	return s
}
