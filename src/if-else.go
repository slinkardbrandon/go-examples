package main

import "fmt"

func main() {

	// There is no ternary if in Go, so we have to use a
	// full if statement even for basic conditions. :(
	const test = 1

	if test == 2 {
		fmt.Println("wtf")
	} else if test == 3 {
		fmt.Println("Seriously.. what?")
	} else {
		fmt.Println("Actually.. it's", test)
	}

}
