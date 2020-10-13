package main

import "fmt"

// type List []int

// func (l List) Len() int { return len(l) }

// func (l *List) Append(val int) { *l = append(*l, val) }
// func (l List) Append(val int) { l = append(l, val) }

func main() {
	// 值
	// var lst List
	// lst.Append(1)
	// fmt.Printf("%v (len: %d)", lst, lst.Len()) // [1] (len: 1)
	// // 指针
	// plst := new(List)
	// plst.Append(2)
	// fmt.Printf("%v (len: %d)", plst, plst.Len()) // &[2] (len: 1)
	s := make([]int, 0)
	s = append(s, 1)
	fmt.Println(s)
}
