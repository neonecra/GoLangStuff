package main

import "fmt"

func main() {
	var x [58]string //adding a number makes a array, arrays are static and have a set length -- [] without a number makes a slice which adjust depending on the content (Golang slice's are more like dynamic arrays)
	fmt.Println(x)
	fmt.Println(len(x)) //len() to get amount of elements
	for i, _ := range x {
		fmt.Print(string(i + 65)) // return the string version of the array index (add 65 for fun so it starts at 'A')
	}
	fmt.Println()
	x[5] = "Test"
	fmt.Println(x[5])
}
