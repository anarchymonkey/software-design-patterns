package flyweights

import "fmt"

type TreeType struct {
	Color   string
	Texture string
	Name    string
}

func (t *TreeType) Draw(config TreeType, x, y int) {
	fmt.Println("Drawing a tree with config & positions", config, x, y)
}
