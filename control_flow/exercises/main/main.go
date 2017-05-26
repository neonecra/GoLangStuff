package main

import "fmt"

func main() {
	fmt.Println("\nExercise 1-1") // START
	fmt.Println("Hello, World!")

	fmt.Println("\nExercise 1-2") // START
	fmt.Println("Hello, my name is Aidan!")

	fmt.Println("\nExercise 1-3") // START
	var input string
	fmt.Print("Whats your name? ")
	fmt.Scan(&input)
	fmt.Println("Hello,", input)

	fmt.Println("\nExercise 1-4") // START
	var smallNum int
	var largeNum int
	fmt.Println("Lets do some maths, give me two numbers.")
	fmt.Print("First number: ")
	fmt.Scan(&smallNum)
	fmt.Print("Second number: ")
	fmt.Scan(&largeNum)
	switch { // switch make sure the larger number is always devided by the smaller one
	case largeNum > smallNum:
		fmt.Println(largeNum, "/", smallNum, "= ", largeNum/smallNum, " with a remainder of ", largeNum%smallNum)
	case largeNum < smallNum:
		fmt.Println(smallNum, "/", largeNum, "= ", smallNum/largeNum, " with a remainder of ", smallNum%largeNum)
	default:
		fmt.Println("These numbers are the same.")
	}

	fmt.Println("\nExercise 1-5") // START
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	fmt.Println("\nExercise 1-6") // START
	for i := 0; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Print("FizzBuzz")
		case i%3 == 0:
			fmt.Print("Fizz")
		case i%5 == 0:
			fmt.Print("Buzz")
		default:
			fmt.Print(i)
		}
		fmt.Print(" ")
	}
	fmt.Println()

	fmt.Println("\nExercise 1-7") // START
	var sum int
	for i := 0; i < 1000; i++ {
		switch {
		case i%3 == 0:
			sum += i
			continue
		case i%5 == 0:
			sum += i
			continue
		default:
			continue
		}
	}
	fmt.Println(sum)

	fmt.Println("\nExercise 2-1") // START
	fmt.Println("\nExercise 2-2") // START
	fmt.Println("\nExercise 2-3") // START
	fmt.Println("\nExercise 2-4") // START
	fmt.Println("\nExercise 2-5") // START
	fmt.Println("\nExercise 2-6") // START

}

/* notes
- For loops (uses break, continue and fallthrough) take a initializer; condition and post this is so the initial value is kept as close as possible , looped through and can be incremented easily, however you can also work it similar to a while stament and only include the condition, or include no conditions and create an endless loop that has to be broken via an internal process.append
- if statements work similar if not the same to javascript, however allow for an initializer to also be included
- a switch statement uses 'case' to match contiditons, if the conditions equate to true then the statement is run, go doesnt have native fallthrough so after one is run the statement ends, however you can define fallthroguh for each subsequent statement where its needed, they work pretty much like a if statement but have that bit of extra added functionality there allowing one condition to flow into another
- runes -- runes (uses ' ') are the characters we use in a language, more than one rune concatinated together is a string (uses " ") and a string using backticks (uses ` `)allows for parragraph intenting and spaces to be retained when they usually arent --- runes are also an alius for int32 because of how it also represents up to 4 bytes
*/
