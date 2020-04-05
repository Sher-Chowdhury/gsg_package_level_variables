package main

import "fmt"

// Here we haven't specified datatype, so golang works it out for you, aka type inference. 
var (
	city = "New York"
	country = "America"
)

func declaredVariablesUsingTypeInference() {
    fmt.Println("##### EG6 output #####")

	// Here we haven't specified datatype, so golang works it out for you, aka type inference. 
	// So here we are actually initialising+declare+infer a variable
	var name = "Peter Parker"

	// there's a special
	// colon-equal syntax, ":=" shorthand you can use instead for initialising+declare+infer a variable. 
    // see https://tour.golang.org/basics/10
	// this only works INSIDE a function
	// this syntax is commonly used. 

	age := 17

	fmt.Println("'name' is set to:", name)  // 'name' is set to: Peter Parker
	fmt.Println("'city' is set to:", city)  // 'city' is set to: New York
	fmt.Println("'age' is set to:", age)    // 'age' is set to: 18
}


