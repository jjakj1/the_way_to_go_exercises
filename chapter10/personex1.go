package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

// call by value 更改的只是原值的一份副本，并不改变原值
func upPerson(per Person) {
	per.firstName = strings.ToUpper(per.firstName)
	per.lastName = strings.ToUpper(per.lastName)
}

func main() {
	// value type
	var per1 Person
	per1.firstName = "Chris"
	per1.lastName = "Wood"
	upPerson(per1)
	fmt.Println(per1)
	// pointer
	per2 := new(Person)
	per2.firstName = "Chris"
	per2.lastName = "Wood"
	upPerson(*per2)
	fmt.Println(*per2)
	// literal
	per3 := &Person{"Chris", "Wood"}
	upPerson(*per3)
	fmt.Println(*per3)
}
