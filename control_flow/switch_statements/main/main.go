package main

import "fmt"

func main() {
	/* switch statements work a little like an if statement sequence but there is more control of what is
	returned since you can control its 'fallthrough', it usually expects at least one outcome to be true,
	but if none are it uses 'default' (in this case it can also be more reliable to
	just use if-else-if-else depending on the situation)
	*/
	switch "dan" {
	case "jenny", "tim": // multiple cases
		fmt.Println("hi jenny, or tim")
		// automatically breaks after first match (because this is the naure of GoLang)
	case "dan":
		fmt.Println("hi dan")
		fallthrough //you can specify fallthrough however to run the next 'case' if there is one (cannot fallthrough the final item in a switch)
	default:
		fmt.Println("you have no friends")
	}

	myFriend := "ma"
	switch {
	case len(myFriend) == 2: // len() gets length of string
		fmt.Println("hello people with name with a length of 2") // returns only the first item that matches
	case myFriend == "ma":
		fmt.Println("hi ma")
	case myFriend == "marcus", myFriend == "tim": // multiple cases
		fmt.Println("hi marcus, or tim")
	case myFriend == "dan":
		fmt.Println("hi dan")
	default:
		fmt.Println("you have no friends")
	}

	fmt.Println(checkType('a')) // run funcation checktype below
}

func checkType(x interface{}) string { //interface{} is for an empty interface since we dont know what it is yet (about 50% sure this is right)
	switch x.(type) { //this is called asserting, to check what type 'x' is
	case int: // is 'x' a int
		return "int"
	case bool: // is 'x' a boolean
		return "bool"
	case int32: // is 'x' a int32
		return "int32"
	case string: // is 'x' a string
		return "string"
	default: // type is unknown
		return "unknown"
	}
}
