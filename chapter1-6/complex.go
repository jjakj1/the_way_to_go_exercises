package main

import "fmt"

func main() {
	c1 := 5 + 10i
	c2 := complex(5, 10)

	fmt.Printf("c1 = %v, c2 = %v\n", c1, c2)
}