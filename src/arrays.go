package main

import "fmt"

func main() {
	fmt.Println("Array example: ")

	// This declares an array with a length of exactly 1 item
	var a [1]int
	// Set that one item (index of 0, woot) to 3.
	a[0] = 3

	fmt.Println(a) // Should see the output [3]

}
