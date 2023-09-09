package main

import (
	"github.com/anarchymonkey/design-patterns/structural/flyweight/components"
	"github.com/anarchymonkey/design-patterns/structural/flyweight/flyweights"
)

func main() {

	treeTypeFactory := flyweights.NewTreeFactory()

	colors := []string{
		"red",
		"green",
		"yellow",
		"orange",
		"red",
		"violet",
		"green",
		"green",
		"green",
		"green",
	}

	names := []string{
		"Hibiscus",
		"Lavendar",
		"Mango",
		"Apple",
		"Hibiscus",
		"Lavendar",
		"Mango",
		"Apple",
		"Hibiscus",
		"Lavendar",
	}

	textures := []string{
		"grainy",
		"grainy",
		"grainy",
		"grainy",
		"grainy",
		"grainy",
		"grainy",
		"grainy",
		"grainy",
		"grainy",
	}

	for i := 0; i < 10; i = i + 1 {
		treeType := treeTypeFactory.GetTreeType(colors[i], names[i], textures[i])
		tree := components.Tree{
			PosX: 55 * i,
			PosY: 100 * i,
		}
		tree.Draw(treeType)
	}
}
