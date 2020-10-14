package main

import (
	f64 "./float64"
	"fmt"
	"sort"
)

func main() {
	f := []float64{3, 2, 4, 5, 1}
	fs := f64.Float64Array(f)
	sort.Sort(fs)
	fmt.Println(fs)

	fs = f64.Fill()
	sort.Sort(fs)
	fmt.Println(fs)
}
