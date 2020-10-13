package stack

import "fmt"

type stack struct {
	arr      []int
	cur, len int // cur 表示下一个元素插入的下标
}

func New(cap int) (s *stack) {
	s = new(stack)
	s.arr = make([]int, cap)
	s.cur, s.len = 0, cap
	return
}

func (s *stack) Push(i int) {
	if s.cur < s.len {
		s.arr[s.cur] = i
		s.cur++
	}
}

func (s *stack) Pop() int {
	if s.cur > 0 {
		tmp := s.arr[s.cur-1]
		s.cur--
		return tmp
	}
	return 0
}

func (s *stack) String() string {
	res := ""
	for i := 0; i < s.cur; i++ {
		res += fmt.Sprintf("[%v:%v]", i, s.arr[i])
	}
	return res
}
