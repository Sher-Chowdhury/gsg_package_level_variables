package main

import (
	"fmt"
)

func initialisedonly() {
	// here we declared to variables but we haven't assigned a value to them.
	// i.e. declared but not yet initialized
	// by default unitialized int-variables takes the value '0' and
	// unitialized string-variables takes the value ''.
	var (
		name string
		city string
		age  int
	)

	fmt.Println("'name' is set to:", name)
	fmt.Println("'city' is set to:", city)
	fmt.Println("'age' is set to:", age)
}
