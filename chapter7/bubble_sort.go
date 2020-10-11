package main

import (
	"fmt"
)

func main() {
	arr := [5]int{2, 3, 1, 5, 4}
	bubbleSort(arr[:])
	fmt.Println(arr)
}

func bubbleSort(arr []int) {
	for i := range arr {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
