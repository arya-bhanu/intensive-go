package main

import (
	"errors"
	"fmt"
	"math"
	"slices"
	"sort"
)

const PHI = math.Pi

func main() {
	fmt.Println("Hello World")
	tri := Triangle{
		SideA: 3,
		SideB: 4,
		SideC: 5,
	}
	circ := Circle{
		Radius: 2,
	}
	fmt.Println(tri.Area())
	fmt.Println(circ.Area())
	fmt.Println(math.Pow(2, 2))
}

type Shape interface {
	Area() float64
	Perimeter() float64
	String() string
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	SideA, SideB, SideC float64
}

func (r *Rectangle) Area() float64 {

	return r.Height * r.Width
}

func (r *Rectangle) Perimeter() float64 {
	return 2*r.Height + 2*r.Width
}

func (r *Rectangle) String() string {
	return fmt.Sprintf("Rectangle %.2f %.2f width height", r.Width, r.Height)
}

func (r *Circle) Area() float64 {
	return PHI * math.Pow(r.Radius, 2)
}

func (r *Circle) Perimeter() float64 {
	return PHI * 2 * r.Radius
}

func (r *Circle) String() string {
	return fmt.Sprintf("Circle %.2f radius", r.Radius)
}

func (r *Triangle) Area() float64 {
	s := (r.SideA + r.SideB + r.SideC) / 2
	area := math.Sqrt(s * (s - r.SideA) * (s - r.SideB) * (s - r.SideC))
	return area
}

func (r *Triangle) Perimeter() float64 {
	return r.SideA + r.SideB + r.SideC
}

func (r *Triangle) String() string {
	return fmt.Sprintf("Triangle %.2f %.2f %.2f sides", r.SideA, r.SideB, r.SideC)
}

func NewRectangle(width, height float64) (*Rectangle, error) {
	if width <= 0.0 || height <= 0.0 {
		return nil, errors.New("width and height could not be less than 0")
	}
	return &Rectangle{
		Width:  width,
		Height: height,
	}, nil
}
func NewCircle(radius float64) (*Circle, error) {
	if radius <= 0.0 {
		return nil, errors.New("radius could not be less than 0")
	}
	return &Circle{
		Radius: radius,
	}, nil
}
func NewTriangle(a, b, c float64) (*Triangle, error) {
	if a <= 0.0 || b <= 0.0 || c <= 0.0 {
		return nil, errors.New("all sides could not be less than 0")
	}

	arrSides := []float64{a, b, c}
	slices.Sort(arrSides)

	if arrSides[0]+arrSides[1] <= arrSides[2] {
		return nil, errors.New("sum of 2 sides must be larger than the other 1 side")
	}

	return &Triangle{
		SideA: a,
		SideB: b,
		SideC: c,
	}, nil
}

type ShapeCalculator struct{}

func NewShapeCalculator() *ShapeCalculator {
	return &ShapeCalculator{}
}
func (sc *ShapeCalculator) PrintProperties(s Shape) {
	fmt.Println(s.String())
}
func (sc *ShapeCalculator) TotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}
func (sc *ShapeCalculator) LargestShape(shapes []Shape) Shape {
	maxShape := struct {
		key   Shape
		value float64
	}{
		key:   shapes[0],
		value: shapes[0].Area(),
	}
	for _, shape := range shapes {
		currArea := shape.Area()
		if maxShape.value < currArea {
			maxShape.key = shape
			maxShape.value = currArea
		}
	}
	return maxShape.key
}
func (sc *ShapeCalculator) SortByArea(shapes []Shape, ascending bool) []Shape {
	sort.Slice(shapes, func(i, j int) bool {
		if ascending {
			return shapes[i].Area() < shapes[j].Area()
		} else {
			return shapes[i].Area() > shapes[j].Area()
		}
	})
	return shapes
}
