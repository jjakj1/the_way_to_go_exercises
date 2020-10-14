package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Shape struct{}

func (s *Shape) Area() float32 {
	return -1
}

type Square struct {
	side float32
	Shape
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

type Rectangle struct {
	width, length float32
	Shape
}

func (r *Rectangle) Area() float32 {
	return r.width * r.length
}

func main() {
	sh := Shape{}
	s := &Square{3, sh}
	r := &Rectangle{3, 4, sh}

	shapes := []Shaper{s, r, &sh}

	for _, shape := range shapes {
		fmt.Println(shape.Area())
	}
}
