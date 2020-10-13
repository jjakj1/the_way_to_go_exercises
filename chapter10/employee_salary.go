package main

import (
	"fmt"
)

type employee struct {
	salary float64
}

func main() {
	em := &employee{12000}
	em.giveRaise(15)
	fmt.Println(em)
}

func (em employee) giveRaise(percent float64) {
	raise := em.salary * (percent / float64(100))
	em.salary += raise
}
