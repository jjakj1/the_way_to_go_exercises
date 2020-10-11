package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8, 9, 10}
	s := insertSlice(s1, s2, 3)
	fmt.Println(s)
}

func insertSlice(s1 []int, s2 []int, pos int) (s []int) {
	n := len(s1)
	m := len(s2)
	s = make([]int, m+n)
	copy(s[0:pos], s1[0:pos])
	copy(s[pos:pos+len(s2)], s2)
	copy(s[pos+len(s2):], s1[pos:])
	return
}
