package main

import (
	"fmt"
)

func initialisedonly() {
	// here we DECLARE to variables but we haven't assigned a value to them.
	// i.e. DECLARED but not yet INITIALISED
	// by default unitialized int-variables takes the value '0' and
	// uninitialized string-variables takes the value ''.
	// i.e. they are initialised behind the scenes be giving them default values.
	var (
		name string
		city string
		age  int
	)

	fmt.Println("'name' is set to:", name)
	fmt.Println("'city' is set to:", city)
	fmt.Println("'age' is set to:", age)
}
