package main

import "fmt"

type Shape interface {
	area() float64
}
type Rectangle struct {
	width, height float64
}

func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

type Circle struct {
	radius float64
}

func (circle Circle) area() float64 {
	return circle.radius * circle.radius * 3.14
}

func main() {
	var shape Shape
	shape = Rectangle{width: 10, height: 5}
	fmt.Printf("Rectangle area is %f\n", shape.area())

	shape = Circle{radius: 5}
	fmt.Printf("Circle area is %f\n", shape.area())
}
