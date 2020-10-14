package main

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	val int
}

func (s Simple) Get() int {
	return s.val
}

func (s Simple) Set(v int) {
	s.val = v
}

type RSimple struct {
	val int
}

func (r RSimple) Get() int {
	return r.val
}

func (r RSimple) Set(v int) {
	r.val = v
}

func fi(sim Simpler) {
	switch sim.(type) {
	case *Simple:
		fmt.Println("This val is of type Simple")
	case *RSimple:
		fmt.Println("This val is of type RSimple")
	default:
		fmt.Println("This type implements Simpler interface")
	}
}

func isSimple(sim Simpler) {
	fmt.Println(sim)
	if t, ok := sim.(*Simple); ok {
		fmt.Println("Yes:", t)
	}
	fmt.Println("No")
}

func main() {
	s1 := &Simple{1}
	s2 := &RSimple{2}

	// fi(s1)
	// fi(s2)

	isSimple(s1)
	isSimple(s2)
}
