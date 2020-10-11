package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	res := Map(multipleByTen, arr[:])
	fmt.Println(res)
}

func Map(f func(int) int, arr []int) []int {
	for i := range arr {
		arr[i] = f(arr[i])
	}
	return arr
}

func multipleByTen(a int) int {
	return a * 10
}
