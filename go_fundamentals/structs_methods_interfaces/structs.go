package structs

import "math"

type Rectangle struct {
	length  float64
	breadth float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.length + rectangle.breadth)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.length * rectangle.breadth
}
