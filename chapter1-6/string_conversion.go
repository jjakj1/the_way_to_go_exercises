package main

import (
	"fmt"
	"strconv"
)

func main() {
	orig := "666L"
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, _ := strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)

	newS := strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
