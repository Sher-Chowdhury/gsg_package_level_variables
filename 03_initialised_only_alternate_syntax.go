package main

import "fmt"

func initialisedonlyAlt2() {

    fmt.Println("##### EG3 output #####")  

	// you can declare multiple variables of the same type on a single line.
	// although this can make the code harder to read
	var (
		name, city string
		age        int
	)

  
	fmt.Println("'name' is set to:", name)   // 'name' is set to:
	fmt.Println("'city' is set to:", city)   // 'city' is set to:
	fmt.Println("'age' is set to:", age)     // 'age' is set to: 0
}