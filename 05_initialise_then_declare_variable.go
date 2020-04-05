package main

import (
	"fmt"
)

func initialiseThenDeclareVariable() {
	fmt.Println("##### EG5 output #####")
	
	var fullname string       // initialised here
	fullname = "Peter Parker" // then declared here

	fmt.Println("'fullname' is set to:", fullname)
}
