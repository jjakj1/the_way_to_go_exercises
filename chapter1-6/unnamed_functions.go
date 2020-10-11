package main

import (
	"fmt"
)

func main() {
	res := func(x, y int) int {
		return x + y
	}(3, 4)
	fmt.Println(res)
}
