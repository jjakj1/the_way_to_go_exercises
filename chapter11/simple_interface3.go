package main

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	val int
}

func (s *Simple) Get() int {
	return s.val
}

func (s *Simple) Set(v int) {
	s.val = v
}

func main() {
	s := &Simple{1}
	gI(s)
}
func gI(any interface{}) {
	if _, ok := any.(Simpler); ok {
		fmt.Println("Is simpler")
		return
	}
	fmt.Println("Not a simpler")
}
