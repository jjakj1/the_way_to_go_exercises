package main

import (
	"fmt"
)

func main() {
	printAllArgs(1, 2, 3, 4, 5)
}

func printAllArgs(s ...int) {
	for _, val := range s {
		fmt.Println(val)
	}
}
