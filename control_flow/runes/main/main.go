package main

import "fmt"

func main() {
	// []byte() prints a 'slice of byte' ie the string of bytes for the given input
	// here we can see that this chinese character is made of 3 bytes while the english character consists of only one, because of how UTF-8 works
	fmt.Println("世 =", []byte("世"))
	fmt.Println("A =", []byte("A"))
	fmt.Println()

	// Characters are also known as runes because of how they reperesent a alius of int32 and UTF-8 is a 4 byte coding scheme (32 bits / 8 = 4 bytes)
	/*	a rune is:
		- character
		- an integer value identifying a Unicode code point
		- alias for int32
			- how many bytes in 32 bits? (4 bytes → 4 * 8 = 32)
			- UTF-8 is a 4 byte coding scheme
			- with int32 (4 bytes) we have numbers for each of the code points
		- a string is a collection of runes
	*/
	for i := 100; i < 150; i++ {
		fmt.Println(i, " | ", []byte(string(i)), " | ", string(i)) // string() is a conversion turning the binary UTF-8 number into a string
	}
	fmt.Println()

	// '  ' single quotes identify a rune, while "  " double quotes identify a string (different from javascript) and if you use `  ` (back-ticks the one on the ~ key) it maintaines spacing and breaks in a string
	// in essence a 'string' is a concatenated version of several 'runes'
	foo := 'a'
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", rune(foo)) //if you convert the int32 item into a rune, its still int32, since 'rune' is an alius for int32 -- and how we store characters/runes
}
