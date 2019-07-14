package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	fmt.Printf("Toying with %v!\n", "methods")

	r := rect{width: 50, height: 100}
	fmt.Printf("Area of rectangle with a width of %v and a height of %v is %v.\n", r.width, r.height, r.area())
	fmt.Printf("Permiter of rectangle with a width of %v and a height of %v is %v.\n", r.width, r.height, r.perim())
}
