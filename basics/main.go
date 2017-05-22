package main

import (
	"fmt"
)

func main() {
	// Print text on new line
	fmt.Println("Hello, World!")

	// Print formated text using formatting verbs -- this line it just to help organise the loop below
	fmt.Printf("%v \t\t %v \t %v \t %v \t\t %v \n", "char", "decimal", "binary", "hex", "utf-8")
	// Print characters between '120' and '128'
	for i := 120; i < 128; i++ {
		fmt.Printf("%c \t\t %d \t\t %b \t %#x \t\t %q \n", i, i, i, i, i)
	}
}
