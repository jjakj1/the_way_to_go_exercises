package main

import (
	"fmt"
)

func main() {
	arr := []float32{1.0, 2.0, 3.0}
	fmt.Println(sum(arr))
}

func sum(arrF []float32) (s float32) {
	for _, value := range arrF {
		s += value
	}
	return
}
