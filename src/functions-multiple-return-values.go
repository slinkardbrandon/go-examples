package main

import "fmt"

func main() {

	fmt.Println(getResult(49))
}

func getResult(a int) (int, int) {
	return a, a * 5
}
