package main

import "fmt"

type NewStringer interface {
	String() string
}

type Base struct {
	val int
}

func (b *Base) String() string {
	return fmt.Sprintf("[%d]", b.val)
}

func (b *Base) Get() int {
	return b.val
}

func main() {
	b := &Base{1}
	var s NewStringer = b
	if sv, ok := s.(NewStringer); ok {
		fmt.Println(sv.Get())
	}
}
