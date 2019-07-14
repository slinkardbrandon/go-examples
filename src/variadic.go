package main

import "fmt"

func main() {
	sum(4, 5, 6, 7)
}

func sum(numbers ...int) {
	total := 0

	for _, num := range numbers {
		fmt.Printf("Calculating: %v + %v = %v\n", total, num, total+num)
		total += num
	}

	fmt.Println("Result: ", total)

}
