package main

// recursion usually not recommended because of performance cost

import "fmt"

func factoral(x int) int {
	if x == 0 {
		return 1
	}
	return x * factoral(x-1) // function calls itself till it equals '0', since it isn't called again at this point the recursion ends
}

func main() {
	fmt.Println(factoral(4))
}
