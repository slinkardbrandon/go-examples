package main

import "fmt"

func main() {
	fmt.Println("Playing with slices")

	// Unlike arrays, slices are not determined by their length.
	// Go slices are probably fairly comparable to typescript arrays in terms of typing (obviously more strict)

	slice := make([]string, 3)

	fmt.Println("Empty slice", slice)

	slice[0] = "test 1"
	slice[1] = "test 2"
	slice[2] = "test 3"

	fmt.Println("Three tests in slice", slice)
	fmt.Printf("Slice length: %v\n", len(slice))

	// Reasonably easy to append to a slice here too. Eseentially array.push() in js
	slice = append(slice, "test 4", "test 5")

	fmt.Println("Five tests in slice", slice)

	// Slices can be copied if there is an empty slice with the same length..
	slice2 := make([]string, len(slice))
	copy(slice2, slice)

	fmt.Println("Copied slice \"slice 2\"", slice2)

	// Kind of neat thing about slices is the "slice" operator, which is essentially
	// inputing a magic string  of "#low:#high" like so slice
	basicSlice := make([]int, 5)
	basicSlice[0] = 0
	basicSlice[1] = 1
	basicSlice[2] = 2
	basicSlice[3] = 3
	basicSlice[4] = 4

	fmt.Println("Slice-n-diced slice", basicSlice[3:5])

	// But we can also do more magic things like this...

	// This slices up to 5, but excludes "5"
	fmt.Println(":5", basicSlice[:5]) // Expected output: [0 1 2 3 4]

	// This slices up from the provided number
	fmt.Println("3:", basicSlice[3:]) // Expected output: [3 4]

}
