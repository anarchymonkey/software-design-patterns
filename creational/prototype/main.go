package main

import (
	"fmt"

	"github.com/anarchymonkey/design-patterns/creational/prototype/shapes"
)

func main() {
	fmt.Println("This is the prototype pattern")

	// slice of clones
	clones := []shapes.Shape{}

	circle := shapes.Circle{}
	rectangle := shapes.Rectangle{}

	circle.Init(10, 20, "red")
	rectangle.Init(20, 30, "blue")

	for i := 0; i < 10; i += 1 {
		if i%2 == 0 {
			clones = append(clones, circle.Clone())
		}
		clones = append(clones, rectangle.Clone())
	}

	for index, clone := range clones {
		fmt.Printf("The clone at index %d is %v with type %T \n", index, clone, clone)
	}

	fmt.Printf("%T", circle)
}
