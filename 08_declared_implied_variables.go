package main

import "fmt"

func declaredImpliedVariables() {

	// This is more shorthand, instead of the using the 'var' word to defined
	// these as variables we have used ':=' to get golang to work that out for us.
	name := "Peter Parker"
	city := "New York"
	age := 18

	fmt.Println("'name' is set to:", name)
	fmt.Println("'city' is set to:", city)
	fmt.Println("'age' is set to:", age)
}
