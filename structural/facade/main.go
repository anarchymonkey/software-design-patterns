package main

import "fmt"

type Shape interface {
	Draw() string
}

type Rectangle struct {
	Width  int
	Height int
}

func (rect *Rectangle) Draw() string {
	var AREA int = rect.Width * rect.Height
	return fmt.Sprintf("Drawing a rectangle with width %d and height %d with area %d", rect.Width, rect.Height, AREA)
}

type Circle struct {
	Radius int
}

func (cir *Circle) Draw() string {
	return fmt.Sprintf("Drawing a circle with radisus %d", cir.Radius)
}

type Triangle struct {
	Base   int
	Height int
}

func (tri *Triangle) Draw() string {
	return fmt.Sprintf("Drawing a triangle with Base %d and Height %d", tri.Base, tri.Height)
}

type Facade struct {
	circle    Circle
	triangle  Triangle
	rectangle Rectangle
}

func (facade *Facade) DrawByType(shapeType string) string {
	switch shapeType {
	case "circle":
		return facade.circle.Draw()
	case "rectangle":
		return facade.rectangle.Draw()
	case "triangle":
		return facade.triangle.Draw()
	default:
		fmt.Println("This shape is not supported")
		return ""
	}
}

func main() {
	facade := Facade{
		circle: Circle{
			Radius: 15,
		},
		triangle: Triangle{
			Base:   10,
			Height: 20,
		},
		rectangle: Rectangle{
			Width:  20,
			Height: 20,
		},
	}

	// facade only exposes the things that client would use and not the internal working
	// internal working can be calculating area before drawing or precomputing coordinates
	fmt.Println(facade.DrawByType("rectangle"))
}
