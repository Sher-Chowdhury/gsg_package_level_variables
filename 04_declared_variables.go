package main

import (
	"fmt"
)


var name string = "Peter Parker"



func declaredVariables() {
    fmt.Println("##### EG4 output #####")

	var (
		city string = "New York"
	)

	// if you are Initialising + Declaring variables, like above, then notice you can do this 
	// either inside or outside of the function.
	fmt.Println("'name' is set to:", name)
	fmt.Println("'city' is set to:", city)
}
