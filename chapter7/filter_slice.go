package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(s)
	newS := filterSlice(s, isOdd)
	fmt.Println(newS)
}

func filterSlice(s []int, f func(int) bool) (res []int) {
	res = []int{}
	for _, v := range s {
		if f(v) {
			res = append(res, v)
		}
	}
	return
}

func isOdd(n int) bool {
	if n%2 == 1 {
		return true
	}
	return false
}
