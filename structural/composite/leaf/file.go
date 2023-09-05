package leaf

import "github.com/anarchymonkey/design-patterns/structural/composite/component"

type Leaf struct {
	Size int
	Name string
}

func (leaf *Leaf) GetSize() int {
	return leaf.Size
}

func (leaf *Leaf) GetName() string {
	return leaf.Name
}

func (leaf *Leaf) SetSize(size int) {
	leaf.Size = size
}

func (leaf *Leaf) SetName(name string) {
	leaf.Name = name
}

func (leaf *Leaf) IsLeafNode() bool {
	return true
}

func (leaf *Leaf) AddFileComponent(fileComponent component.FileComponent) {
	panic("This method is not supported, adding a file component is not supported")
}

func (leaf *Leaf) PrintSubComponents() {
	panic("This method is not suported, printing file components is not suported")
}
