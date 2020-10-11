package main

import (
	"fmt"
)

func main() {
	sFrom := []int{1, 2, 3}
	sTo := make([]int, 10)
	n := copy(sTo, sFrom) // 将后者的内容复制到前者中，返回复制的元素个数
	fmt.Printf("copied %d elements\n", n)
	fmt.Println(sTo)

	// 添加单个元素
	sTo = append(sTo, 4, 5, 6)
	fmt.Println(sTo) // [1 2 3 0 0 0 0 0 0 0 4 5 6] 直接添加到容量后面
	// 添加一个切片
	sAppend := []int{7, 8, 9}
	sTo = append(sTo, sAppend...)
	fmt.Println(sTo)
}
