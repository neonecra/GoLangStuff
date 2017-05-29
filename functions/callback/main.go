package main

import (
	"fmt"
)

/* NOTE - -not always recommended to use callbacks (because not very idiomatic and take a little longer to read), but it is useful to know how to read them */

/* function with no receiver but has parameters and a return
func | identifier | parameters                       | return */
func filter(numbers []int, evaluate func(int) bool) []int {
	/* 	numbers []int			-- gives a variable to slice of ints
	   	evaluate func(int) bool	-- identifies a callback, the function as named 'evaluate', this function takes an 'int' and returns a bool
	   	[]int 					-- this function returns a slice of int */
	var xs []int
	for _, i := range numbers { // loop through the items in the 'numbers' slice, first return is the index, second is the value (we discard the index since we dont need it)
		if evaluate(i) {
			xs = append(xs, i)
		}
	}
	return xs
}

func main() {
	// callbacks are not super readable so if theres a more idiomatic way of coding it then do it that way :)
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool { /* pass a slice of int and a function which will be called back as part of the 'filter' function */
		return n > 1
	})
	fmt.Println(xs) // [2,3,4]
}
