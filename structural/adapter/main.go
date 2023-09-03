package main

import (
	"fmt"

	"github.com/anarchymonkey/design-patterns/structural/adapter/adapter"
	"github.com/anarchymonkey/design-patterns/structural/adapter/entities"
)

func main() {
	fmt.Println("This is the main function")

	roundHole := entities.RoundHole{
		Radius: 10,
	}

	fmt.Println("The round hole width is", roundHole.GetRadius())

	roundPeg := entities.RoundPeg{
		Radius: 9,
	}

	fmt.Println("Does the roundpeg fit in the round hole", roundHole.Fits(roundPeg))

	squarePeg := entities.SquarePeg{
		Width: 12,
	}

	squarePegAdapter := adapter.SquarePegAdapter{
		SquarePeg: squarePeg,
	}

	fmt.Println("Does the squarePeg fit in the roundHole", roundHole.Fits(squarePegAdapter.GetRoundPeg()))

}
