package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(sum(arr[:]))

}

// 计算数组元素和
func sum(a []int) (s int) {
	for i := range a {
		s += a[i]
	}
	return
}
