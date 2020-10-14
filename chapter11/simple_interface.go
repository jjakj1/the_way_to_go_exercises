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
	fmt.Println(s.Get())
	s.Set(2)
	fmt.Println(s.Get())
}
