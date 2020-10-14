package stack

// import "fmt"

type Stack struct {
	Data []interface{}
	Last int
}

func (s *Stack) Len() int {
	return len(s.Data)
}

func (s *Stack) IsEmpty() bool {
	return s.Last == 0
}

func (s *Stack) Push(x interface{}) {
	s.Data[s.Last] = x
	s.Last++
}

func (s *Stack) Pop() interface{} {
	s.Last--
	return s.Data[s.Last]
}

// func (s *Stack) String() string {
// 	res := "["
// 	for i := 0; i < s.Last; i++ {
// 		res += fmt.Sprintf("%d")
// 		if i != s.Last-1 {
// 			res += ", "
// 		}
// 	}
// 	res += "]"
// 	return res
// }
