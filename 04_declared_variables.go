package main

import (
	"fmt"
)

func declaredVariables() {

	var name string = "Peter Parker"

	var (
		city string = "New York"
	)

	// if you are Initialising + Declaring variables, like above,  then you do this using colon-equal syntax, ":="
	// instead, it is more short-hand and more common:

	age := 17

	fmt.Println("'name' is set to:", name)
	fmt.Println("'city' is set to:", city)
	fmt.Println("'age' is set to:", age)
}
