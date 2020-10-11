package main

import (
	"fmt"
)

func main() {
	items := [5]int{10, 20, 30, 40, 50}
	for index := range items {
		items[index] *= 2
	}
	fmt.Println(items)
}
