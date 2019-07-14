package main

import "fmt"

func main() {
	fmt.Println(plus(1, 5))
	fmt.Println(plus(6, 9))
}

func plus(a int, b int) int {
	return a + b
}
