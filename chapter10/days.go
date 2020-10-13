package main

import (
	"fmt"
)

type Day int

const (
	Mon Day = iota
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func main() {
	var th Day = 3
	fmt.Println(th)
}

func (day Day) String() string {
	return dayName[day]
}
