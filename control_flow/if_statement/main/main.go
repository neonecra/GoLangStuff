package main

import "fmt"

func main() {
	if false {
		fmt.Println("false")
	} else if err := 404; err == 404 { // can initialise with a variable like in for statements, helps keep code tight and idiomatic
		fmt.Println("oh noh! 404")
	}
}
