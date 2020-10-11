package main

import (
	"fmt"
)

func main() {
	printNum(10)
}

func printNum(start int) {
	if start <= 0 {
		return
	}
	fmt.Println(start)
	printNum(start - 1)
}
