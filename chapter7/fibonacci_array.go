package main

import (
	"fmt"
)

func main() {
	var fiboArray [50]int
	for i := range fiboArray {
		if i <= 1 {
			fiboArray[i] = 1
		} else {
			fiboArray[i] = fiboArray[i-1] + fiboArray[i-2]
		}
	}

	for i, v := range fiboArray {
		fmt.Println(i, ":", v)
	}
}
