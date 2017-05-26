package main

import (
	"fmt"
)

func main() {
	fmt.Print(shout("Dan ", "Smith "))
}

func shout(fname string, lname string) (string, string) { // you can also return multiple items by comma delimiting the returned values and defining their types
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname) // fmt.Sprint() concatinates strings together
}
