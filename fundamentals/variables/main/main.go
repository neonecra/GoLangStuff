package main

import (
	"fmt"

	"github.com/neonecra/GoLangStuff/variables/stringutil"
)

func main() {
	// grab and print variable from a package
	fmt.Println(stringutil.MyName)
	fmt.Println()

	// shorthand variable (only within 'func') <- best method
	a := "dragon"
	fmt.Println(a)
	fmt.Println()

	// variable set to zero value then assigned (only within 'func')
	var b int
	b = 10
	fmt.Printf("%T - %v \n\n", b, b)

	// zero value of variables
	var c int
	var d string
	var e bool
	var f float64
	fmt.Printf("%T - %v \n", c, c)
	fmt.Printf("%T - %v \n", d, d)
	fmt.Printf("%T - %v \n", e, e)
	fmt.Printf("%T - %v \n", f, f)
}
