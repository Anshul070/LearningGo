package main

import (
	"math"
	"fmt"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

// 2. Define concrete types
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Triangle struct {
	Hypotenuse float64
	Height float64
	Base float64
}

func (t Triangle) Area () float64 {
	a := (t.Base * t.Height) / 2
	return a
}

func (t Triangle) Perimeter() float64 {
	p := t.Height + t.Base + t.Hypotenuse
	return p
}

// 3. Generic function for Shape interface
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f ", s.Area())
	fmt.Printf("Perimeter: %.2f ", s.Perimeter())
	fmt.Println()
}

func main(){
	shapes := []Shape{
		Rectangle{3, 4},
		Circle{2.5},
		Triangle{15, 5, 8},
	}

	fmt.Println("Iterating over multiple shapes:")
	for i, s := range shapes {
		fmt.Printf("\nShape %d: %T\n", i+1, s)
		printShapeInfo(s)
	}
}
