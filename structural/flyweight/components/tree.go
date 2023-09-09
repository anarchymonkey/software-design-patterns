package components

import "github.com/anarchymonkey/design-patterns/structural/flyweight/flyweights"

type Tree struct {
	PosX    int
	PosY    int
	TreeMap flyweights.TreeType
}

func (tree *Tree) Draw(config flyweights.TreeType) {
	tree.TreeMap.Draw(config, tree.PosX, tree.PosY)
}
