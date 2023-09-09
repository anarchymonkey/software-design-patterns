package flyweights

import "fmt"

type TreeFactory struct {
	treeTypes map[string]TreeType
}

func NewTreeFactory() *TreeFactory {
	return &TreeFactory{
		treeTypes: make(map[string]TreeType),
	}
}

func (treeFact *TreeFactory) GetTreeType(color, texture, name string) TreeType {
	key := fmt.Sprintf("%s_%s_%s", color, name, texture)
	treeType, ok := treeFact.treeTypes[key]

	if !ok {
		newTreeType := TreeType{
			Color:   color,
			Name:    name,
			Texture: texture,
		}
		treeFact.treeTypes[key] = newTreeType
		return newTreeType
	}
	fmt.Println("returning from cache", treeType)
	return treeType
}
