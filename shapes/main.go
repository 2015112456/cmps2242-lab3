package main

import "fmt"

func main() {
	r := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Created rectangle: %v\n", r)
	FindAreaPerimeter(r)

	c := Circle{Radius: 8}
	fmt.Printf("\nCreated circle: %v\n", c)
	FindAreaPerimeter(c)

	t := Triangle{Base: 4, Height: 5}
	fmt.Printf("\nCreated triangle: %v\n", t)
	FindAreaPerimeter(t)

	r.Scale(3)
	fmt.Printf("\nAfter scaling rectangle by 3: %v\n", r)
	FindAreaPerimeter(r)
}

func FindAreaPerimeter(shape any) {
	area, perimeter := 0.0, 0.0

	switch s := shape.(type) {
	case Rectangle:
		area = s.Area()
		perimeter = s.Perimeter()
	case Circle:
		area = s.Area()
		perimeter = s.Perimeter()
	case Triangle:
		area = s.Area()
		perimeter = s.Perimeter()
	}
	fmt.Printf("Area = %.2f\n", area)
	fmt.Printf("Perimeter = %.2f\n", perimeter)
}
