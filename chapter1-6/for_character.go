package main

import (
	"fmt"
)

func main() {
	// for i := 0; i < 25; i++ {
	// 	for j := 0; j < i; j++ {
	// 		fmt.Print("G")
	// 	}
	// 	fmt.Println()
	// }
	for i, str := 0, "G"; i < 25; i, str = i+1, str+"G" {
		fmt.Println(str)
	}
}
