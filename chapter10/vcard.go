package main

import (
	"fmt"
	"time"
)

type Address struct {
	country string
	city    string
	street  string
}

type VCard struct {
	name    string
	address *Address
	birth   time.Time
	phote   string
}

func main() {
	vc := &VCard{
		"lsy",
		&Address{"China", "Beijing", "HaiDian"},
		time.Date(1990, 7, 29, 0, 0, 0, 0, time.Local),
		"C:/111.jpg",
	}

	fmt.Printf("the full vcard is: %v", vc)

}
