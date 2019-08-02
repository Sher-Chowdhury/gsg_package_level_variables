package main

import (
	"fmt"
)

func initialiseThenDeclareVariable() {

	var fullname string       // initialised here
	fullname = "Peter Parker" // then declared here

	fmt.Println("'fullname' is set to:", fullname)
}
