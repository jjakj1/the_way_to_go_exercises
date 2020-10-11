package main

import (
	"fmt"
)

func main() {
	fv := func() {
		fmt.Println("Hello world!")
	}

	fmt.Printf("the type of fv is %T and the value of fv is %p\n", fv, fv)
}
