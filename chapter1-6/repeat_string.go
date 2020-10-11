package main

import (
	"fmt"
	"strings"
)

func main() {
	origS := "Hi there"
	newS := strings.Repeat(origS, 3)
	fmt.Println("The new repeated string is:", newS)
}
