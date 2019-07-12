package main

import "fmt"

func main() {
	// For is go's only looping construct, no do/while, no map, etc

	fmt.Println("Basic loop")
	fmt.Println()

	// Most basic for loop example:
	i := 0
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("More shorthand")
	fmt.Println()
	// Let's short hand it a tiny bit
	for x := 0; x < 3; x++ {
		fmt.Println(x)
	}
}
