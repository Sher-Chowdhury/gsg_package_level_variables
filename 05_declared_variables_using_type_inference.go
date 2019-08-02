package main

import "fmt"

func declaredVariablesUsingTypeInference() {

	var (
		name = "Peter Parker"
		city = "New York"
		age  = 18
	)

	fmt.Println("'name' is set to:", name)
	fmt.Println("'city' is set to:", city)
	fmt.Println("'age' is set to:", age)
}
