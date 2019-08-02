package main

import "fmt"

func main() {

	fmt.Println("EG1 - Initialised only")
	initialisedonly()

	fmt.Println("EG2 - Initialised only - Alternate syntax")
	initialisedonlyAlt1()

	fmt.Println("EG3 - Initialised only - Alternate syntax")
	initialisedonlyAlt2()

	fmt.Println("EG4 - Declared Variables")
	declaredVariables()

	fmt.Println("EG5 - Initialise then declare")
	initialiseThenDeclareVariable()

	fmt.Println("EG6 - Declared Variables using type inference")

	declaredVariablesUsingTypeInference()

	fmt.Println("EG7 - Declared Variables - Alternate syntax")
	declaredVariablesAlt1()

	fmt.Println("EG8 - Check data type")
	checkDataType()
}
