package main

import (
	"fmt"
)

func main() {
	b := [5]byte{'a', 'b', 'c', 'd', 'e'}
	first, second := splitSlice(b[:], 3)
	fmt.Println(first)
	fmt.Println(second)
}

func splitSlice(buffer []byte, n int) (first, second []byte) {
	return buffer[:n], buffer[n:]
}
