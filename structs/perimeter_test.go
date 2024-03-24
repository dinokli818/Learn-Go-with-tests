package structs

import (
	"testing"
)

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, hasArea float64) {
		t.Helper()
		got := shape.Area()
		if got != hasArea {
			t.Errorf("%v got: %g but hasAreaed: %g", shape, got, hasArea)
		}
	}

	areaTests := []struct {
		shape   Shape
		hasArea float64
		name    string
	}{
		{shape: Rectangle{Width: 10.0, Height: 4.0}, hasArea: 40.0, name: "Rectangle"},
		{shape: Circle{Redius: 10}, hasArea: 314.1592653589793, name: "Circle"},
		{shape: Triangle{Width: 12, Height: 6}, hasArea: 36.0, name: "Triangle"},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.hasArea)
		})

	}

}
