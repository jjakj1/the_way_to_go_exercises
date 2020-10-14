package main

import (
	"fmt"
)

func main() {
	num := 1
	str := "abc"
	fmt.Println(mapFunc(num))
	fmt.Println(mapFunc(str))
}

func mapFunc(data interface{}) interface{} {
	if i, ok := data.(int); ok {
		return i * 2
	}
	if i, ok := data.(string); ok {
		return i + i
	}
	return nil
}
