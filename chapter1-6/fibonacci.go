package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		result, pos := fibonacci(i)
		fmt.Printf("fibonacci(%d) is %d\n", pos, result)
	}
}

func fibonacci(n int) (res int, pos int) {
	if n <= 1 {
		res, pos = 1, n
	} else {
		res1, _ := fibonacci(n - 1)
		res2, _ := fibonacci(n - 2)
		res = res1 + res2
		pos = n
	}
	return
}
