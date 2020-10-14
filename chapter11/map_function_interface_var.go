package main

import (
	"fmt"
)

func main() {
	num := 1
	str := "abc"
	mapFunc(num, str)
}

func mapFunc(data ...interface{}) {
	for i := range data {
		switch data[i].(type) {
		case int:
			fmt.Println(data[i].(int) * 2)
		case string:
			fmt.Println(data[i].(string) + data[i].(string))
		default:
		}
	}
}
