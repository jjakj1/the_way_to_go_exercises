package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func Abs(p *Point) float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type Polar struct {
	x float64
	y float64
	z float64
}

func Scale(p *Polar, fac float64) {
	p.x *= fac
	p.y *= fac
	p.z *= fac
}

func main() {
	poi := &Point{3.0, 4.0}
	pol := &Polar{3.0, 4.0, 5.0}
	fmt.Println(Abs(poi))
	Scale(pol, 2.0)
	fmt.Println(pol)
}
