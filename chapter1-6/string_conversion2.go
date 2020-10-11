package main

import (
	"fmt"
	"strconv"
)

func main() {
	orig := "ABC"
	_, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	}
}
