package main

import "fmt"

const test string = "test"

const (
	_  = iota             // blank identifier to dump first result
	KB = 1 << (iota * 10) // '<<' is bit-shift, so moves to 'tens' place in bits ( i.e = 2^10 )
	MB                    // the previous constat value increments with the 'iota' value going up by 1 every time (bit-shift 2^20)
	GB                    // (bit-shift 2^20)
)

func main() {
	fmt.Printf("%v \n", test)
	fmt.Printf("%v %b \n", KB, KB)
	fmt.Printf("%v %b \n", MB, MB)
	fmt.Printf("%v %b \n", GB, GB)
}
