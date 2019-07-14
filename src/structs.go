package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Structs!")

	fmt.Println(person{name: "Bob", age: 42})

	bob := person{name: "Bob", age: 30}
	// Access properties of a struct with `.`
	fmt.Printf("Bob's name: %v.\n", bob.name)
	fmt.Printf("Bob's age: %v.\n", bob.age)

	// Cannot access via index though
	// fmt.Printf("Bob's age: %v.\n", bob[age])

	// Structs are mutable, same as JS
	bob.age = 52
	fmt.Printf("Bob's new age: %v.\n", bob.age)

}
