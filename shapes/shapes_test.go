package main

import (
	"math"
	"testing"
)

// helper function to use tolerance for float comparison
func almostEqual(a, b float64) bool {
	const epsilon = 0.01
	return math.Abs(a-b) < epsilon
}

func TestRectangleArea(t *testing.T) {
	tests := []struct {
		name string
		r    Rectangle
		want float64
	}{
		{name: "Area of Rectangle (3, 4)", r: Rectangle{Width: 3, Height: 4}, want: 12.0},
		{name: "Area of Rectangle (3.3, 4.4)", r: Rectangle{Width: 3.3, Height: 4.4}, want: 14.52},
		{name: "Area of Rectangle (333.3, 4.4)", r: Rectangle{Width: 333.3, Height: 4.4}, want: 1466.52},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r.Area()
			if !almostEqual(got, tt.want) {
				t.Errorf("got %.2f, want %.2f", got, tt.want)
			}
		})
	}
}

func TestRectanglePerimeter(t *testing.T) {
	tests := []struct {
		name string
		r    Rectangle
		want float64
	}{
		{name: "Perimeter of Rectangle (10, 5)", r: Rectangle{Width: 10, Height: 5}, want: 30.0},
		{name: "Perimeter of Rectangle (10.1, 5.5)", r: Rectangle{Width: 10.1, Height: 5.5}, want: 31.2},
		{name: "Perimeter of Rectangle (100.1, 5.5)", r: Rectangle{Width: 100.1, Height: 5.5}, want: 211.2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r.Perimeter()
			if !almostEqual(got, tt.want) {
				t.Errorf("got %.2f, want %.2f", got, tt.want)
			}
		})
	}
}

func TestCircleArea(t *testing.T) {
	tests := []struct {
		name string
		c    Circle
		want float64
	}{
		{name: "Area of Circle (3)", c: Circle{Radius: 3}, want: 28.27},
		{name: "Area of Circle (3.3)", c: Circle{Radius: 3.3}, want: 34.21},
		{name: "Area of Circle (50)", c: Circle{Radius: 50}, want: 7853.98},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.Area()
			if !almostEqual(got, tt.want) {
				t.Errorf("got %.2f, want %.2f", got, tt.want)
			}
		})
	}
}

func TestCirclePerimeter(t *testing.T) {
	tests := []struct {
		name string
		c    Circle
		want float64
	}{
		{name: "Perimeter of Circle (3)", c: Circle{Radius: 3}, want: 18.85},
		{name: "Perimeter of Circle (3.3)", c: Circle{Radius: 33.3}, want: 209.23},
		{name: "Perimeter of Circle (50)", c: Circle{Radius: 50}, want: 314.16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.Perimeter()
			if !almostEqual(got, tt.want) {
				t.Errorf("got %.2f, want %.2f", got, tt.want)
			}
		})
	}
}

func TestTriangleArea(t *testing.T) {
	tests := []struct {
		name string
		tr   Triangle
		want float64
	}{
		{name: "Area of Triangle (3, 4)", tr: Triangle{Base: 3, Height: 4}, want: 6.0},
		{name: "Area of Triangle (5.5, 6.6)", tr: Triangle{Base: 5, Height: 6}, want: 15.0},
		{name: "Area of Triangle (100, 8)", tr: Triangle{Base: 100, Height: 8}, want: 400.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tr.Area()
			if !almostEqual(got, tt.want) {
				t.Errorf("got %.2f, want %.2f", got, tt.want)
			}
		})
	}
}

func TestTrianglePerimeter(t *testing.T) {
	tests := []struct {
		name string
		tr   Triangle
		want float64
	}{
		{name: "Perimeter of Triangle (3, 4)", tr: Triangle{Base: 3, Height: 4}, want: 9.0},
		{name: "Perimeter of Triangle (5.5, 6.6)", tr: Triangle{Base: 5.5, Height: 6.6}, want: 16.5},
		{name: "Perimeter of Triangle (100, 8)", tr: Triangle{Base: 100, Height: 8}, want: 300},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tr.Perimeter()
			if !almostEqual(got, tt.want) {
				t.Errorf("got %.2f, want %.2f", got, tt.want)
			}
		})
	}
}

func TestRectangleScale(t *testing.T) {
	tests := []struct {
		name   string
		r      Rectangle
		factor float64
		want   Rectangle
	}{
		{name: "Scale Rectangle (2, 3) by 2", r: Rectangle{Width: 2, Height: 3}, factor: 2, want: Rectangle{Width: 4, Height: 6}},
		{name: "Scale Rectangle (5, 5) by 0.5", r: Rectangle{Width: 5, Height: 5}, factor: 0.5, want: Rectangle{Width: 2.5, Height: 2.5}},
		{name: "Scale Rectangle (10, 1) by 3", r: Rectangle{Width: 10, Height: 1}, factor: 3, want: Rectangle{Width: 30, Height: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r
			got.Scale(tt.factor)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircleScale(t *testing.T) {
	tests := []struct {
		name   string
		c      Circle
		factor float64
		want   Circle
	}{
		{name: "Scale Circle (2) by 2", c: Circle{Radius: 2}, factor: 2, want: Circle{Radius: 4}},
		{name: "Scale Circle (5) by 0.5", c: Circle{Radius: 5}, factor: 0.5, want: Circle{Radius: 2.5}},
		{name: "Scale Circle (10) by 3", c: Circle{Radius: 10}, factor: 3, want: Circle{Radius: 30}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c
			got.Scale(tt.factor)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTriangleScale(t *testing.T) {
	tests := []struct {
		name   string
		tr     Triangle
		factor float64
		want   Triangle
	}{
		{name: "Scale Triangle (2, 3) by 2", tr: Triangle{Base: 2, Height: 3}, factor: 2, want: Triangle{Base: 4, Height: 6}},
		{name: "Scale Triangle (5, 5) by 0.5", tr: Triangle{Base: 5, Height: 5}, factor: 0.5, want: Triangle{Base: 2.5, Height: 2.5}},
		{name: "Scale Triangle (10, 1) by 3", tr: Triangle{Base: 10, Height: 1}, factor: 3, want: Triangle{Base: 30, Height: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tr
			got.Scale(tt.factor)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
