package tests

import (
	"goland-demo/demo/shapes"
	"testing"
)

func TestArea(t *testing.T) {

	areaTests := []struct {
		name     string
		shape    shapes.Shape
		expected float64
	}{
		{name: "Rectangle", shape: shapes.Rectangle{Width: 12, Height: 6}, expected: 72.0},
		{name: "Circle", shape: shapes.Circle{Radius: 10}, expected: 314.1592653589793},
		{name: "Triangle", shape: shapes.Triangle{Base: 12, Height: 6}, expected: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.expected {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.expected)
			}
		})
	}

}
