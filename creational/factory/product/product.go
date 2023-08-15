package product

import (
	"errors"
	"fmt"
)

type TransportVehicle interface {
	Deliver(from, to, message string) (string, error)
}

type Truck struct {
	Code string
}

type Ship struct {
	Code string
}

func (truck Truck) Deliver(from, to, message string) (string, error) {
	if from == "Canada" && to == "India" {
		return "", errors.New("cannot transfer to other countries")
	}
	return fmt.Sprintf("Delivering from %s to %s with message %s", from, to, message), nil
}

func (ship Ship) Deliver(from, to, message string) (string, error) {
	if from == "Canada" && to == "India" {
		return "", errors.New("cannot transfer to other countries")
	}
	return "transport started from lat 52 & long 102", nil
}
