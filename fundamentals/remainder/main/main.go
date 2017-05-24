package main

import "fmt"

func main() {
	x := 13 / 3 //while '/' devides like usua;
	y := 13 % 3 //using '%' operator gets the remainder
	fmt.Printf("13 / 4 = %v with a remainder of %v", x, y)

	for i := 0; i < 10; i++ {
		if (i % 2) == 1 {
			fmt.Println("odd")
		} else {
			fmt.Println("even")
		}
	}
}
