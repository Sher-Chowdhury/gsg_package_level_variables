package main

import (
	"fmt"
)

func declaredVariables() {

	var name string = "Peter Parker"

	var (
		city string = "New York"
		age  int    = 18
	)

	fmt.Println("'name' is set to:", name)
	fmt.Println("'city' is set to:", city)
	fmt.Println("'age' is set to:", age)
}
