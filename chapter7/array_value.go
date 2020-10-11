package main

import (
	"fmt"
)

func main() {
	var arr [5]int
	fmt.Printf("%p\n", &arr)
	arr[1] = 100
	fmt.Printf("%p\n", &arr)
}
