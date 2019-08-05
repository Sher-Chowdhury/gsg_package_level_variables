package main

import "fmt"

func thePointerDatatype() {
	// A pointer is a variable which points to the memory location of another variable.
	// you can view memory location using '&'
	fullname := "Peter Parker"
	fmt.Println("memory location for 'fullname' variable is:", &fullname)

	// This initialises a pointer, which just points to memory location.
	// the rhs also ends up declaring this and set allocates the memory.
	var firstname *string = new(string)

	// To stor anything inside the memory, you use asterisk
	*firstname = "peter"

	// this prints out the memory location
	fmt.Println(firstname)
	// this prints out what's stored in the memory location itself.
	fmt.Println(*firstname)

}
