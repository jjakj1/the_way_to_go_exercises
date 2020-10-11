package main

import (
	"fmt"
	"sort"
)

func main() {
	drinks := map[string]string{
		"cola":   "可乐",
		"spirit": "雪碧",
		"wine":   "红酒",
		"beer":   "啤酒",
	}

	for key := range drinks {
		fmt.Printf("%v | ", drinks[key])
	}

	fmt.Println()

	keys := make([]string, 4)
	i := 0
	for key, value := range drinks {
		fmt.Printf("%v: %v | ", key, value)
		keys[i] = key
		i++
	}

	fmt.Println()

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%v: %v | ", key, drinks[key])
	}
}
