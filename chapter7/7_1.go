package main

import (
	"fmt"
)

func main() {
	// a := [...]string{"a", "b", "c", "d"}
	// for i, _ := range a {
	// 	fmt.Println("Array item", i, "is", a[i])
	// }

	arr1 := new([3]int)
	arr2 := *arr1
	arr2[2] = 100
	fmt.Println(*arr1)
	fmt.Println(arr2)
}
