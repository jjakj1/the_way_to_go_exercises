package main

import (
	"fmt"
	"sort"
)

type Person struct {
	firstName, lastName string
}

type Persons []Person

func main() {
	ps := Persons{Person{"George", "Washington"}, Person{"John", "Adams"}, Person{"Thomas", "Jefferson"}}
	// ps := Persons(p)
	sort.Sort(ps)
	fmt.Println(ps)
}

func (ps Persons) Len() int {
	return len(ps)
}

func (ps Persons) Less(i, j int) bool {
	return ps[i].firstName < ps[j].firstName || ps[i].lastName < ps[j].lastName
}

func (ps Persons) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
