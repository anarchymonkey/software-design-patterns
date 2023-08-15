package creator

import (
	"errors"

	"github.com/anarchymonkey/design-patterns/creational/factory/product"
)

type TransportType string

const (
	TRUCK      TransportType = "TRUCK"
	TRAIN      TransportType = "TRAIN"
	SHIP       TransportType = "SHIP"
	SUB_MARINE TransportType = "SUB_MARINE"
)

// concrete creator 1
func RoadTransportFactory(transportType TransportType) (product.TransportVehicle, error) {
	switch transportType {
	case TRUCK:
		return product.Truck{
			Code: "1102",
		}, nil
	default:
		return nil, errors.New("this type is not supported")
	}
}

// concrete creator 2
func SeaTransportFactory(transportType TransportType) (product.TransportVehicle, error) {
	switch transportType {
	case SHIP:
		return product.Ship{
			Code: "MS5102",
		}, nil
	default:
		return nil, errors.New("ship type not supported")
	}
}
