package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4}
	fmt.Println(len(slice), cap(slice))
	newSlice := extendSlice(slice, 3)
	fmt.Println(len(slice), cap(slice))
	fmt.Println(len(newSlice), cap(newSlice))
}

func extendSlice(s []int, factory int) []int {
	if factory <= 0 {
		return s
	}
	newS := make([]int, factory*len(s))
	copy(newS, s)
	s = newS
	return s
}
