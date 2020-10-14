package float64

import (
	"fmt"
	"math/rand"
)

type Float64Array []float64

func (f Float64Array) Len() int {
	return len(f)
}

func (f Float64Array) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f Float64Array) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f Float64Array) List() string {
	res := "["
	for i := 0; i < len(f); i++ {
		res += fmt.Sprintf("%v, ", f[i])
	}
	res += "]"
	return res
}

func (f Float64Array) String() string {
	return f.List()
}

func NewFloat64Array() Float64Array {
	f := make([]float64, 25)
	fs := Float64Array(f)
	return fs
}

func Fill() Float64Array {
	f := make([]float64, 10)
	for i := range f {
		f[i] = rand.Float64()
	}
	fs := Float64Array(f)
	return fs
}
