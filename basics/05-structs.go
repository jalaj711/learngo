package main

import "fmt"

type car struct {
	make             string
	name             string
	ground_clearance int
	wheelType        wheel
}

type wheel struct {
	radius   int
	material string
}

// nested anonymous struct
type car2 struct {
	make             string
	name             string
	ground_clearance int
	wheelType        struct {
		radius   int
		material string
	}
}

// embeded structs
type truck struct {
	car
	capacity int
}

// methods on structs
type rect struct {
	width  int
	length int
}

func (r rect) area() int {
	return r.length * r.width
}

func main() {
	mycar := car{
		make:             "abcd",
		name:             "name",
		ground_clearance: 56,
		wheelType: wheel{
			radius:   4,
			material: "rc4",
		},
	}
	fmt.Printf("Car make is %s\n", mycar.make)

	truck1 := truck{
		car: car{
			name:             "abcvd",
			make:             "ejvrbu",
			ground_clearance: 56,
			wheelType: wheel{
				radius:   4,
				material: "rc4",
			},
		},
		capacity: 40,
	}

	// point to note is that properties "inherited" from car are
	// accessible at top level of truck
	fmt.Printf("Truck capacity is %d and make is %s\n", truck1.capacity, truck1.make)

	r := rect{length: 10, width: 5}
	fmt.Printf("Area of the rectangle is %d\n", r.area())
}
