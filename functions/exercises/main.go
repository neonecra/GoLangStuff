package main

import "fmt"

func main() {

	// Write a function which takes an integer. The function will have two returns. The first return should be the
	// argument divided by 2. The second return should be a bool that let’s us know whether or not the argument
	// was even. For example:
	// - half(1) returns (0, false)
	// - half(2) returns (1, true)
	fmt.Println("\nExercise 2-1")

	myIntegerThing := func(x int) (int, bool) {
		return x / 2, x%2 == 0
	}
	half, even := myIntegerThing(2)
	fmt.Println(half, even)

	// Modify the previous program to use a func expression.
	fmt.Println("\nExercise 2-2")

	highest := func(x ...int) int {
		var winner int
		for _, z := range x {
			if winner < z {
				winner = z
			}
		}
		return winner
	}
	fmt.Println(highest(1, 4, 2, 6, 2, 0, 9, 2, 5))

	// Write a function with one variadic parameter that finds the greatest number in a list of numbers.
	fmt.Println("\nExercise 2-3")

	// What's the value of this expression: (true && false) || (false && true) || !(false && false)?
	fmt.Println("\nExercise 2-4")
	fmt.Println((true && false) || (false && true) || !(false && false))

	// Write a function, foo, which can be called in all of these ways:
	fmt.Println("\nExercise 2-5")
	var foo = func(x ...int) {
		var sum int
		for _, y := range x {
			sum += y
		}
		fmt.Println(sum)
	}
	var example = func() {
		foo(1, 2)
		foo(1, 2, 3)
		aSlice := []int{1, 2, 3, 4}
		foo(aSlice...)
		foo()
	}
	example()

	// Find a problem at projecteuler.net then create the solution. Add a comment beneath your solution that includes a description of the problem. Upload your solution to github. Tweet me a link to your solution.
	fmt.Println("\nExercise 2-6")
	fibCount := func(x int, y int) {
		if y+x == 0 {
			x := 1
			y := 2
		}
		if 4000000 < x {
			return
		}
		fmt.Println(x + y) // 3
		if x > y {
			fibCount(y, x+y)
		}
	}

}
