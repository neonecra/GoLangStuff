package stringutil

import "fmt"

var MyName = "potato"
var futureName = MyName + " and he's super awesome"

func ProcessName(x string) {
	fmt.Println(x + futureName)
}

func PrintThis(x string) {
	fmt.Println(x)
}

func PrintThisNumber(x int) {
	fmt.Println(x)
}
