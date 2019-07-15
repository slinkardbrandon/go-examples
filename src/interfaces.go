package main

import "fmt"
import "math"

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return r.width*2 + r.height*2
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(shape geometry) {
	fmt.Println(shape)
	fmt.Println("area:", shape.area())
	fmt.Println("perimeter:", shape.perimeter())
}

func main() {
	fmt.Printf("Toying around with %v!\n", "interfaces")

	c := circle{radius: 60}
	r := rectangle{width: 81, height: 93}

	fmt.Println("\nMeasuring circle")
	measure(c)

	fmt.Println("\nMeasuring rectangle")
	measure(r)

}
