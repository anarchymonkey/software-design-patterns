package creator

import (
	"fmt"
)

type Transport interface {
	Deliver()
}

type Truck struct {
	ratePerHour int
}

type Ship struct {
	ratePerHour int
}

func (truck Truck) Deliver() {
	fmt.Println("Truck will deliver your goods in 1-3 biz days, rate per hour is", truck.ratePerHour)
}

func (ship Ship) Deliver() {
	fmt.Println("Ship will deliver your goods in 1-3 biz days, rate per hour is", ship.ratePerHour)
}

// this is a factory pattern, we are generating the data dynamically with a particular transport type, so the deliver method remains the same but the implementation differs by a huge margin
func TransportFactory(transportType string, rate int) Transport {

	switch transportType {
	case "Truck":
		return &Truck{ratePerHour: rate}
	case "Ship":
		return &Ship{ratePerHour: rate}
	default:
		panic("Wront transporttype selected")
	}
}
