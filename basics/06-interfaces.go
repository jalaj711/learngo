package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
	// named parameters and returns
	smallerThanShape(shapeToBeCompared shape) (isSmaller bool)
}

type rect struct {
	width  int
	length int
}

func (r rect) area() float64 {
	return float64(r.width * r.length)
}

func (r rect) perimeter() float64 {
	return float64(2 * (r.length + r.width))
}

func (r rect) smallerThanShape(shapeToBeCompared shape) bool {
	return float64(r.width*r.length) < shapeToBeCompared.area()
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return float64(3.14159 * c.radius * c.radius)
}

func (c circle) perimeter() float64 {
	return float64(2 * 3.14159 * c.radius)
}

func (c circle) smallerThanShape(shapeToBeCompared shape) bool {
	return float64(3.14159*c.radius*c.radius) < shapeToBeCompared.area()
}

func is_area_equal(s1, s2 shape) bool {
	return s1.area() == s2.area()
}

// type assertion with interfaces
func is_a_rect(s shape) bool {
	_, ok := s.(rect)
	if ok {
		return true
	}
	return false
}

// type assertion with interfaces
// but with type switches
func type_of_shape(s shape) string {
	switch s.(type) {
	case rect:
		return "Rectangle"
	case circle:
		return "Circle"
	default:
		return "Unknown"
	}
}

func main() {
	r := rect{
		width:  12,
		length: 10,
	}
	c := circle{
		radius: 2.3534,
	}

	fmt.Printf("Areas are equal? ")
	if is_area_equal(r, c) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	fmt.Printf("Is c a rect? ")
	if is_a_rect(c) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	fmt.Printf("Type of r is: %s\n", type_of_shape(r))

}
