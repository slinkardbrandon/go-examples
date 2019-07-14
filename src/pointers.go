package main

import "fmt"

func main() {
	fmt.Println("Pointers :D")

	i := 1
	zeroval(i)

	fmt.Printf("Value after zeroing out the provided val: \"%v\".\n", i)

	zeropointer(&i)

	fmt.Printf("Value after zeroing out the provided pointer: \"%v\".\n", i)

}

// This will create a new reference in memory for "val" (essentially creating a copy), any modifications done to "val"
// will not reflect on the provided value but will be modified to this memory reference
func zeroval(val int) {
	val = 0
}

// This will accept a reference to an int value that already exists in memory.
// Any modifications to "pointer" will reflect on the provided value"
// The syntax *int means that this takes an "int" pointer. *string would accept a string pointer in a similar way
func zeropointer(pointer *int) {
	*pointer = 0
}
