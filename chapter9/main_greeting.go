package main

import (
	"./greetings/morning"
	"./greetings/night"
	"fmt"
)

func main() {
	name := "lsy"
	goodDay := morning.GoodDay(name)
	fmt.Println(goodDay)

	goodNight := night.GoodNight(name)
	fmt.Println(goodNight)
}
