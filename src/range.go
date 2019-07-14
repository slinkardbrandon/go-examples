package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	sum := 0

	fmt.Println(nums)

	// For each number
	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)
}
