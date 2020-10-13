package main

import (
	"fmt"
)

type newS struct {
	f float32
	int
	string
}

func main() {
	s := &newS{1.2, 1, "hello"}
	fmt.Println(s)
}
