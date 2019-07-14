package main

import "fmt"

func main() {
	fmt.Println("Playing with maps now")

	// Looks like they're declared the same way as slices are (using the builtin make command). Neat.
	// "The make built-in function allocates and initializes an object of type slice, map, or chan (only)."

	// This declaration is essentially declaring a map with a key type of "string" but a value type of "int"
	m := make(map[string]int)

	// Print empty map
	fmt.Println(m)

	// Add a couple key/value pairs
	m["dope"] = 1
	m["nasty"] = 2

	fmt.Println(m)

	// Print the value that exists on "nasty" (2)
	fmt.Println(m["nasty"])

}
