package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	s = removeSlice(s, 2, 5)
	fmt.Println(s)
}

func removeSlice(s []int, start, end int) []int {
	m := len(s)
	copy(s[start:], s[end:len(s)])
	s = s[:m-(end-start)]
	return s
}
