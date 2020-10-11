package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	ret, err := mysqrt(-1)
	if err != nil {
		fmt.Println("Error! Return values are:", ret, err)
	} else {
		fmt.Println("It's ok! Return values are:", ret, err)
	}
}

func mysqrt(f float64) (res float64, err error) {
	if f < 0 {
		res, err = float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number")
	} else {
		res, err = math.Sqrt(f), nil
	}
	return
}
