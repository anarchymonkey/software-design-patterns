package adapter

import (
	"math"

	"github.com/anarchymonkey/design-patterns/structural/adapter/entities"
)

type SquarePegAdapter struct {
	SquarePeg entities.SquarePeg
}

func (sqPegAdapter *SquarePegAdapter) GetRoundPeg() entities.RoundPeg {
	return entities.RoundPeg{
		Radius: (sqPegAdapter.SquarePeg.Width * math.Sqrt2) / 2,
	}
}
