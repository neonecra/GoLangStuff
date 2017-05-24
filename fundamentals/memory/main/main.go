package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println()
	fmt.Printf("memory address: %v \n", &a)
	fmt.Printf("decimal name of address: %d \n", &a) //adding '&' before a variable shows its memory location

	var b *int = &a // using '*' operator tells the variable to point to the memory address where the 'int' is stored
	fmt.Print("pointer address: ")
	fmt.Print(b) // 'b' is now equal to that address
	fmt.Println()

	*b = 40 // "using 'b' which is pointing to 'a' we can change the value of 'a'"
	fmt.Print("value at 'a' now is: ")
	fmt.Print(a) // using the '*' operator infront of the variable now gets the contents of that address
	fmt.Println()

	fmt.Println()

	// little thing to tell you about your age
	name := myInput()
	age := myCalculator(name)
	describe := describeAge(age)
	fmt.Print(name + " is ")
	fmt.Print(age)
	fmt.Print(" years " + describe + " \n")
}

func myInput() string {
	var enteredName string
	fmt.Print("Name:")
	fmt.Scan(&enteredName) // fmt.Scan() allows for input and stores in a momory location
	fmt.Println("Name is " + enteredName + "\n")
	return enteredName
}

func myCalculator(x string) float64 {
	var enteredCalculation float64
	fmt.Print(x + "'s age:")
	fmt.Scan(&enteredCalculation)
	return enteredCalculation
}

func describeAge(x float64) string {
	if x <= 10 {
		return "young"
	} else if x <= 20 {
		return "bleh"
	}
	return "old"
}
