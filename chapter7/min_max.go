package main

import (
	"fmt"
)

func main() {
	s := []int{5, 2, 3, 1, 4}
	fmt.Println(minSlice(s))
	fmt.Println(maxSlice(s))
}

func minSlice(s []int) (min int) {
	if len(s) == 0 {
		return 0
	}
	min = s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return
}

func maxSlice(s []int) (max int) {
	if len(s) == 0 {
		return 0
	}
	max = s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return
}
