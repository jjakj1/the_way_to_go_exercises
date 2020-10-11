package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fibo(40)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("the fibonacci calculation consuming time: %s\n", delta)
}

func fibo(pos int) int {
	if pos <= 1 {
		return 1
	}

	return fibo(pos-1) + fibo(pos-2)
}
