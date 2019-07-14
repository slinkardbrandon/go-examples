package main

import "fmt"

func main() {
	sum(4, 5, 6, 7)
}

func sum(numbers ...int) {
	total := 0

	for _, num := range numbers {
		fmt.Println("Calculating: %v + %v", total, num)
		total += num
	}

	fmt.Println("Result: ", total)

}
