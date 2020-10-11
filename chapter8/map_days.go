package main

import (
	"fmt"
)

func main() {
	days := map[string]int{
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6,
		"Sunday":    7,
	}

	if _, ok := days["Tuesday"]; ok {
		fmt.Println("Tuesday exists")
	}

	if _, ok := days["Holleday"]; ok {
		fmt.Println("Holleday exists")
	}

}
