package furnitures

import (
	"errors"
	"fmt"
)

type CoffeeTable interface {
	SitOn()
	HasLegs() (bool, error)
}

type VictorianCoffeeTable struct {
	NoOfLegs int
}

type MordernCoffeeTable struct {
	NoOfLegs int
}

func (victorian *VictorianCoffeeTable) SitOn() {
	fmt.Println("We are sitting on a victorian style chair with legs", victorian.NoOfLegs)
}

func (mordern *MordernCoffeeTable) SitOn() {
	fmt.Println("We are sitting on a mordern chair with legs", mordern.NoOfLegs)
}

func (victorian *VictorianCoffeeTable) HasLegs() (bool, error) {
	if victorian.NoOfLegs < 0 {
		return false, errors.New("furniture should not have negative number of legs")
	}

	return victorian.NoOfLegs != 0, nil
}

func (mordern *MordernCoffeeTable) HasLegs() (bool, error) {
	if mordern.NoOfLegs < 0 {
		return false, errors.New("furniture should not have negative number of legs")
	}

	return mordern.NoOfLegs != 0, nil
}
