package main

import "fmt"

func main() {
	// This declaration is essentially the same as
	thing := "test"
	// Both this
	var thing2 = "test"
	// and this
	var thing3 string = "test"

	// Since both "thing" and "thing2" do not have a strict type set, they default to "string"

	fmt.Println(thing + " 1")
	fmt.Println(thing2 + " 2")
	fmt.Println(thing3 + " 3")
}
