package main

import (
	"fmt"
)

func main() {
	f := fabonacci()
	for i := 0; i < 9; i++ {
		fmt.Println(i+2, f())
	}
}

func fabonacci() func() int {
	first, second := 1, 1
	return func() int {
		first, second = second, first+second
		return second
	}
}
