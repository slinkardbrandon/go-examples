package main

import "fmt"

func main() {
	const hobbitsLocation string = "Isengard"

	switch hobbitsLocation {
	case "Mordor", "Isengard":
		fmt.Printf("They're taking the hobbits to %s!\n", hobbitsLocation)
	default:
		fmt.Println("The Hobbits are just fine.")
	}

}
