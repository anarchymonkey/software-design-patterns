package furnitures

import (
	"errors"
	"fmt"
)

type Chair interface {
	SitOn()
	HasLegs() (bool, error)
}

type VictorianChair struct {
	NoOfLegs int
}

type MordernChair struct {
	NoOfLegs int
}

func (victorian *VictorianChair) SitOn() {
	fmt.Println("We are sitting on a victorian style chair with legs", victorian.NoOfLegs)
}

func (mordern *MordernChair) SitOn() {
	fmt.Println("We are sitting on a mordern chair with legs", mordern.NoOfLegs)
}

func (victorian *VictorianChair) HasLegs() (bool, error) {
	if victorian.NoOfLegs < 0 {
		return false, errors.New("furniture should not have negative number of legs")
	}

	return victorian.NoOfLegs != 0, nil
}

func (mordern *MordernChair) HasLegs() (bool, error) {
	if mordern.NoOfLegs < 0 {
		return false, errors.New("furniture should not have negative number of legs")
	}

	return mordern.NoOfLegs != 0, nil
}
