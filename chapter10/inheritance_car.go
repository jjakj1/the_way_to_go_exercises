package main

import (
	"fmt"
)

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
	wheelCount int
}

type Mercedes struct {
	Car
}

func main() {
	me := new(Mercedes)
	me.wheelCount = 4
	me.Start()
	me.Stop()
	fmt.Println(me.numberOfWheels())
	me.sayHiToMerkel()
}

func (c Car) Stop() {
	fmt.Println("the car stops")
}

func (c Car) Start() {
	fmt.Println("the car starts")
}

func (c Car) numberOfWheels() int {
	return c.wheelCount
}

func (m Mercedes) sayHiToMerkel() {
	fmt.Println("Hi")
}
