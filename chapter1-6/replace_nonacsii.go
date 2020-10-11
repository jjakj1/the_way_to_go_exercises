package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Map(replace, "Jérôme Österreich"))
}

func replace(input rune) rune {
	if input < 127 {
		return input
	} else {
		return ' '
	}
}
