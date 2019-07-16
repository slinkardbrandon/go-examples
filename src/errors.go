package main

import "errors"
import "fmt"

// In go, it's pretty common that errors are the second return value of a function
func somethingThatErrors(i int) (int, error) {
	if i < 5 {
		return i, nil
	}
	return -1, errors.New("test")
}

func main() {
	result, error := somethingThatErrors(4)

	fmt.Println(result, error)

	result, error = somethingThatErrors(8)
	fmt.Println(result, error)

}
