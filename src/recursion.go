package main

import "fmt"

func main() {
	recurse(5)
	recurse(1)
}

func recurse(a int) {
	fmt.Println("recurse! ", a)
	if a < 10 || a > 50 {
		a++
		recurse(a)
	}
}
