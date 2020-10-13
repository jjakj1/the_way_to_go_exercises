package main

import (
	"fmt"
)

type Celsius float64

func main() {
	ce := new(Celsius)
	*ce = Celsius(15.3)
	fmt.Println(ce)
}

func (ce *Celsius) String() string {
	return fmt.Sprintf("%f C", float64(*ce))
}
