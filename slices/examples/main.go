package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5) // 'make' a slice, with the length of '0' and a capacity of '5'
	// mySlice := []int{1,2,3,4,5} making a slice like this will create a slice with a length of the elements and a matching capacity

	fmt.Println(len(mySlice)) // length
	fmt.Println(cap(mySlice)) // capacity
	fmt.Println("------------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, 1) // capacity of slice doubles once its capacity is reached -- but this creates a performance hit
		fmt.Println("len:", len(mySlice), "  capacity:", cap(mySlice))
	}

	fmt.Println("------------------")
	records := make([][]string, 0)
	person1 := []string{
		"rooster",
		"cat",
		"football",
	}
	person2 := []string{
		"pooster",
		"pat",
		"pootball",
	}
	records = append(records, person1, person2)
	fmt.Println(records)
}
