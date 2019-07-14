package main

import "fmt"

func main() {
	fmt.Println("Time for playing with closures")

	var b = bump()
	fmt.Println(b()) // 1
	fmt.Println(b()) // 2
	fmt.Println(b()) // 3

	fmt.Println(bump()()) // 1
}

// Function that returns a function that returns an int
func bump() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
