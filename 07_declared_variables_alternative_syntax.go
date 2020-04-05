package main

import (
	"fmt"
)

func declaredVariablesAlt1() {
	fmt.Println("##### EG7 output #####")
	// here multiple variables of the same data type can be declared on the same line
	// although it makes it a little harder to read
	var (
		name, city = "Peter Parker", "New York"
		age        = 18
	)

	fmt.Println("'name' is set to:", name)
	fmt.Println("'city' is set to:", city)
	fmt.Println("'age' is set to:", age)
	fmt.Println("'age' is set to:", country) // This comes from the eg6 function. This works even if you don't call 
	                                         // eg6 function from main.go
}
