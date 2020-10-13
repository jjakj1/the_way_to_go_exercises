package main

import (
	"fmt"
)

type Rectangle struct {
	height, width int
}

func Area(r *Rectangle) int {
	return r.height * r.width
}

func Perimeter(r *Rectangle) int {
	return 2 * (r.height + r.width)
}

func main() {
	r := &Rectangle{3, 5}
	fmt.Println(Area(r))
	fmt.Println(Perimeter(r))
}
