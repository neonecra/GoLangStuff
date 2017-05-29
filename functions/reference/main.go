package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // []
	changeFirst(m)
	fmt.Println(m) // [Aidan] -- because slices are already referencing a memory address there is no reason to declare '&m' to pass the memory address, all slices of the same origin will be the same
}

func changeFirst(z []string) {
	z[0] = "Aidan"
	fmt.Println(z) // [Aidan]
}
