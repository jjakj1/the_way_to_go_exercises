package main

import (
	"fmt"
	"time"
)

var first, second int

func main() {
	start := time.Now()
	fib(40)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("the fibonacci calculation consuming time: %s\n", delta)
}

func fib(n int) int {
	first, second = 1, 1
	if n <= 1 {
		return 1
	}

	for i := 2; i <= n; i++ {
		first, second = second, first+second
	}
	return second
}
