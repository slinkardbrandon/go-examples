package main

import "fmt"

func main() {
	fmt.Println(`fuck ` + `that`) // String concatenation
	fmt.Println(.01 + .02)        // .03 (floating point accuracy??? :O)
	fmt.Println(true && false)    // False takes precedence
	fmt.Println(true || false)    // true succeeds
	fmt.Println(!true)            // negate values
}
