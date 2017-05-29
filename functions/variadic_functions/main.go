package main

import "fmt"

func main() {
	fmt.Println(average(5, 10, 20, 5))

	//we can also pass a 'slice' of float64 and tell GoLang that its a variadic argument
	list := []float64{5, 5, 5, 5, 30, 30, 30, 30}
	n := average(list...) // by placing the ' ... ' at the end it tells the compiler to pass each part of the slice individually -- OUTGOING variadic
	fmt.Println(n)
}

func average(args ...float64) float64 { // using ' ... ' is variadic, can be used on the last parameter to process an infinite amount of args -- INCOMING variadic
	var total float64
	for _, i := range args {
		total += i
	}
	return total / float64(len(args)) //len() to get amount of arguments to calculate an average
}
