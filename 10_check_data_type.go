package main

// we use the reflect module for finding out a variable's datatype.
import (
	"fmt"
	"reflect"
)

func checkDataType() {
    fmt.Println("##### EG10 output #####")

	/* 
	    the colon-equal syntax, ':=' is the shorthand alternattive to using the 'var' key word to Initialise a variable.
	    this syntax is more common than using the var keyword.
	    You can only use ':=' inside functions
	*/
	
	name := "Peter Parker"
	city := "New York"
	age := 18
	weight := 65.23
	single := true

	// here we find out what data type golang has determined for each variable
	fmt.Println("Golang has decided to set the datatype for 'name' to:", reflect.TypeOf(name))
	fmt.Println("Golang has decided to set the datatype for 'city' to:", reflect.TypeOf(city))
	fmt.Println("Golang has decided to set the datatype for 'age' to:", reflect.TypeOf(age))
	fmt.Println("Golang has decided to set the datatype for 'weight' to:", reflect.TypeOf(weight))
	fmt.Println("Golang has decided to set the datatype for 'single' to:", reflect.TypeOf(single))
}
