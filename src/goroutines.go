package main

import "fmt"

func main() {
	fmt.Printf("Goroutine time baby!\n")

	fmt.Println("Trying goroutine")
	go f("Goroutine")

	f("main")
	fmt.Scanln()
	fmt.Println("Done!")
}

func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
	}
}
