package main

import (
	"fmt"

	"github.com/anarchymonkey/design-patterns/structural/composite/component"
	"github.com/anarchymonkey/design-patterns/structural/composite/composite"
	"github.com/anarchymonkey/design-patterns/structural/composite/leaf"
)

func main() {
	fmt.Println("This is the composite pattern")

	var fileComponent component.FileComponent

	folder := composite.Folder{
		Size: 5040,
		Name: "root",
	}

	leafComponent1 := leaf.Leaf{
		Size: 1250,
		Name: "ext",
	}
	fileComponent = &leafComponent1
	folder.AddFileComponent(fileComponent)

	leafComponent2 := leaf.Leaf{
		Size: 2400,
		Name: "hosts",
	}

	fileComponent = &leafComponent2
	folder.AddFileComponent(fileComponent)

	folder2 := composite.Folder{
		Size: 7045,
		Name: "users",
	}

	leafComponent3 := leaf.Leaf{
		Size: 2040,
		Name: "Desktop",
	}

	leafComponent4 := leaf.Leaf{
		Size: 5040,
		Name: "Documents",
	}

	folder2.AddFileComponent(&leafComponent3)
	folder2.AddFileComponent(&leafComponent4)

	folder.AddFileComponent(&folder2)

	folder.PrintSubComponents()

}
