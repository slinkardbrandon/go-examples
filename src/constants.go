package main

import "fmt"

func main() {
	const num int = 5000000
	const str string = "testing string constants"

	fmt.Println(num)
	fmt.Println(str)

	// A numeric constant has no type until it’s given one, such as by an explicit conversion.
	const nonTypedNum = 5923
	fmt.Println(int64(nonTypedNum))
	fmt.Println(float32(nonTypedNum))

}
