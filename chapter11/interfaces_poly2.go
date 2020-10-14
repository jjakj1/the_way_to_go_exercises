package main

import (
	"fmt"
	"math"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	width, length float32
}

func (re *Rectangle) Area() float32 {
	return re.width * re.length
}

type Circle struct {
	radius float32
}

func (ci *Circle) Area() float32 {
	return math.Pi * ci.radius * ci.radius
}

func main() {
	sq := &Square{3}
	re := &Rectangle{3, 5}
	ci := &Circle{3}

	sh := []Shaper{sq, re, ci}
	for _, s := range sh {
		fmt.Println(s.Area())
	}
}
