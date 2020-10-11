package main

import (
	"./even"
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if even.IsEven(i) {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
