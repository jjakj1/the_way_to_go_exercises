package main

import (
	"fmt"
)

func main() {
	a1, b1, c1 := returnMultiple1(10, 7)
	fmt.Println(a1, b1, c1)
	a2, b2, c2 := returnMultiple2(10, 7)
	fmt.Println(a2, b2, c2)
}

func returnMultiple1(a int, b int) (int, int, int) {
	return a + b, a * b, a - b
}

func returnMultiple2(a int, b int) (sum int, mult int, minus int) {
	sum = a + b
	mult = a * b
	minus = a - b
	return // 命名返回值返回时直接使用一条不带参数的 return 语句
}
