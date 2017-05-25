package main

import "fmt"

func main() {
	// for loop with 'init';'condition';'post'
	for i := 0; i < 5; i++ {
		// nested loop
		for p := 0; p <= i; p++ {
			fmt.Println(i, "-", p)
		}
	}
	fmt.Println()

	// for loop that acts like a 'while' statement -- has to evaluate to a boolean true/false
	var count int
	for count < 10 {
		fmt.Println(count)
		count++
	}
	fmt.Println()

	// endless loop, needs 'break' (or maybe return) to stop
	var endless int
	for {
		if endless >= 9000 { //runs till it gets to 9000, then prints the number
			fmt.Println("Its over", endless, "!")
			break
		}
		endless++
	}
	fmt.Println()

	// endless loop, that prints odd numbers
	var odd int
	for {
		odd++
		if odd%2 == 0 { //when no remainder (i.e. even) end this loop and start the next one
			continue
		}
		fmt.Println(odd)
		if odd >= 50 { //runs till it gets to 50
			break
		}
	}
}
