package main

import "fmt"

func main() {

	m := make(map[string]int) // map base value is nil, when defining 'map[string]' defines the key type, while 'int' is the type of value
	m["k2"] = 8               // define a key and its value at the same time
	m["cat"] = 3
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(m["root"]) // calling key returns 0 when unassigned
	fmt.Println()

	_, ok := m["k2"] // gets bool of map item's assigned status
	fmt.Println(ok)  // true = assigned
	fmt.Println()

	delete(m, "k2")

	_, ok = m["k2"]
	fmt.Println(ok) // false = unassigned
}
