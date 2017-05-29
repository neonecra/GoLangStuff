package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}
func world() {
	fmt.Println("World")
}
func main() {
	defer world() //defer delays a function from running until it BEFORE the parent function closes, usually used to close a file that has been opened to conserve resources
	hello()
}
