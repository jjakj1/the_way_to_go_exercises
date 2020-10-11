package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "The white cat fox is so cut"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str2 := strings.Join(sl, "|")
	fmt.Printf("sl joined by |: %s\n", str2)
}
