package main

import (
	"fmt"
)

type Base struct {
	id int
}

type Person struct {
	Base
	FirstName string
	LastName  string
}

type Employee struct {
	Person
	salary int
}

func main() {
	em := &Employee{Person{Base{1}, "Bruce", "Wayne"}, 1000}
	fmt.Println(em.Id())
}

func (b Base) Id() int {
	return b.id
}

func (b *Base) SetId(i int) {
	b.id = i
}
