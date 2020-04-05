package main

import "fmt"

func declaredImpliedVariablesAlt1() {
	fmt.Println("##### EG9 output #####")
	
	// This is even more short-hand
	name, city, age := "Peter Parker", "New York", 18

	fmt.Println("'name' is set to:", name)
	fmt.Println("'city' is set to:", city)
	fmt.Println("'age' is set to:", age)
	// here's another way of using this
	fmt.Println(name, city, age)

}
