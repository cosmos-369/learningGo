package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f ", got, want)
	}
}

type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{length: 10, breadth: 20}, want: 20},
		{name: "Circle", shape: Circle{radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{base: 10, height: 5}, want: 25},
	}

	for _, tt := range areaTests {

		t.Run(tt.name, func(t *testing.T) {

			got := tt.shape.Area()

			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt, got, tt.want)
			}
		})
	}
}
