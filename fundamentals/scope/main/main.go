package main

import (
	"github.com/neonecra/GoLangStuff/variables/stringutil"
)

// universal variable / package level scope
var universalVar = "I am "

func main() {
	// declaring a test value for an example
	myCaveEcho := "Test"
	stringutil.PrintThis(myCaveEcho)

	// an inner scope, variables within not accessable outside
	{
		// variable within this func scope, processed by another func
		myCaveEcho = caveEcho("Hello")
		// use imported package to process my strings, this function imports fmt so not needed within our top level executable
		stringutil.PrintThis(myCaveEcho + " " + universalVar + "Introducing... ")
	}
	// similar to javascript, re-declaring this updates the original variable
	stringutil.PrintThis(myCaveEcho)

	// call wrapper function, assigns the returned function to 'incremant'
	increment := wrapper()
	stringutil.PrintThisNumber(increment())
	stringutil.PrintThisNumber(increment())
	stringutil.PrintThisNumber(increment())

}

// this function ADD's an enclosed function to whatever calls it
func wrapper() func() int {
	// variables are recursive like JS so this isnt accessable by earlier items in func, this also narrows the scope of the variable (always keep variable near its function if possible)
	z := 0
	// func expression (assigned anonymous function to a variable)
	return func() int {
		z++
		return z
	}
}

func caveEcho(x string) string {
	return x + " " + x + " " + x
}
