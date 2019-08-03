package main

import "fmt"

func main() {

	fmt.Println("EG1 - Initialised only")
	initialisedonly()

	fmt.Println("\nEG2 - Initialised only - Alternate syntax")
	initialisedonlyAlt1()

	fmt.Println("\nEG3 - Initialised only - Alternate syntax")
	initialisedonlyAlt2()

	fmt.Println("\nEG4 - Declared Variables")
	declaredVariables()

	fmt.Println("\nEG5 - Initialise then declare")
	initialiseThenDeclareVariable()

	fmt.Println("\nEG6 - Declared Variables using type inference")

	declaredVariablesUsingTypeInference()

	fmt.Println("\nEG7 - Declared Variables - Alternate syntax")
	declaredVariablesAlt1()

	fmt.Println("\nEG8 - Check data type")
	checkDataType()
}
