package main

import (
	"./min"
	"fmt"
)

func main() {
	ia := min.IntArray{3, 4, 1, 2, 5}
	fmt.Println(min.Min(ia))
}
