package main

import (
	"fmt"
)

func main() {
	s := "google"
	res := copySame([]rune(s))
	fmt.Println(string(res))
}

func copySame(b []rune) []rune {
	res := []rune{}
	for i := range b {
		if i > 0 && b[i] != b[i-1] {
			res = append(res, b[i])
		}
	}
	return res
}
