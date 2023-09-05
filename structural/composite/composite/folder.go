package composite

import (
	"fmt"

	"github.com/anarchymonkey/design-patterns/structural/composite/component"
)

type Folder struct {
	Size              int
	Name              string
	SubFileComponents []component.FileComponent
}

func (folder *Folder) GetSize() int {
	return folder.Size
}

func (folder *Folder) GetName() string {
	return folder.Name
}

func (folder *Folder) SetName(name string) {
	folder.Name = name
}

func (folder *Folder) SetSize(size int) {
	folder.Size = size
}

func (folder *Folder) IsLeafNode() bool {
	return false
}

func (folder *Folder) AddFileComponent(fileComponent component.FileComponent) {
	folder.SubFileComponents = append(folder.SubFileComponents, fileComponent)
}

func (folder *Folder) PrintSubComponents() {
	fmt.Println(folder.GetName())
	for _, folder := range folder.SubFileComponents {
		if !folder.IsLeafNode() {
			folder.PrintSubComponents()
			return
		}
		fmt.Println("_____", folder.GetName(), "size is ->", folder.GetSize())
	}
}
