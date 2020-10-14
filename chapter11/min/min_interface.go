package min

type Miner interface {
	Len() int
	ElemIx(ix int) interface{}
	Less(i, j int) bool
}

func Min(data Miner) interface{} {
	min := data.ElemIx(0)
	minIndex := 0
	for i := 0; i < data.Len(); i++ {
		if data.Less(i, minIndex) {
			min = data.ElemIx(i)
			minIndex = i
		}
	}
	return min
}

type IntArray []int

func (ia IntArray) Len() int                  { return len(ia) }
func (ia IntArray) ElemIx(ix int) interface{} { return ia[ix] }
func (ia IntArray) Less(i, j int) bool        { return ia[i] < ia[j] }
