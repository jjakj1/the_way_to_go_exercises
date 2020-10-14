package main

import (
	"fmt"
	"math"
)

type Magnitude interface {
	Abs() float64
}

type Point struct {
	a, b float64
}

type Solar struct {
	a, b, c float64
}

func main() {
	p := &Point{1, 2}
	s := &Solar{1, 2, 3}
	mags := []Magnitude{p, s}

	for _, m := range mags {
		fmt.Println(m.Abs())
	}
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.a*p.a + p.b*p.b)
}

func (s *Solar) Abs() float64 {
	return math.Sqrt(s.a*s.a + s.b*s.b + s.c*s.c)
}
