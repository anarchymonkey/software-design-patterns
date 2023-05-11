package main

import (
	"fmt"
)

type Passenger struct {
	Name            string
	SeatNo          int
	FlyingMachineId string
}

type FlyingMachine interface {
	Fly(source, destination string, passengers []Passenger)
}

type Airport struct {
	NoOfFlights int
	IsOpen      bool
}

type Airplane struct {
	Id string
}
type Helecopter struct {
	Id string
}

// accepts a flyingMachine type object
func (airport *Airport) Accept(flyingMachine FlyingMachine) {
	// init passengers
	var passengers []Passenger = []Passenger{
		{
			Name:            "Aniket",
			SeatNo:          21,
			FlyingMachineId: "JE-251",
		},
		{
			Name:            "Swarnim",
			SeatNo:          22,
			FlyingMachineId: "FC-620",
		},
	}

	flyingMachine.Fly("CCU", "BLR", passengers)
}

// fly airplane
func (airplane *Airplane) Fly(source, destination string, passengers []Passenger) {
	fmt.Printf("This is an airplane -> flying from source %s and the destination is %s & the passengers are %v \n", source, destination, passengers)
}

// fly helecopter
func (helecopter *Helecopter) Fly(source, destination string, passengers []Passenger) {
	fmt.Printf("This is a helecopter -> flying from source %s and the destination %s & the passengers are %v \n", source, destination, passengers)
}

func main() {

	// create the interface array
	var flyingMachines []FlyingMachine

	// airport has an accept method which acceptes flying maching
	var airport Airport = Airport{
		NoOfFlights: 21,
		IsOpen:      true,
	}

	var airplane Airplane = Airplane{
		Id: "JE-251",
	}
	var helecopter Helecopter = Helecopter{
		Id: "FC-620",
	}

	var flyingMachine FlyingMachine

	// airplane uses flying machine
	flyingMachine = &airplane
	flyingMachines = append(flyingMachines, flyingMachine)

	// helecopter uses flying machine
	flyingMachine = &helecopter
	flyingMachines = append(flyingMachines, flyingMachine)

	/*


		In this example, the Airport struct has an Accept method that accepts a FlyingMachine interface object.
		The FlyingMachine interface defines a method Fly that takes source, destination, and passengers as input parameters. Both Airplane and Helecopter implement the Fly method of the FlyingMachine interface.

		The Airplane and Helecopter types are not exposed outside of the package.
		Instead, the outside world interacts with these types through the Airport struct's Accept method, which encapsulates the Fly method of the FlyingMachine interface.
		The Accept method doesn't reveal the implementation details of the Airplane or Helecopter types but only exposes the functionality of the FlyingMachine interface.
		This way, changes to the implementation of Airplane or Helecopter won't affect the clients that use the Accept method.

		In summary,
		this example demonstrates encapsulation by hiding the implementation details of the Airplane and Helecopter types and providing a well-defined interface (FlyingMachine) for interacting with these types through the Airport struct's Accept method.

	*/

	for _, flyingMachine := range flyingMachines {
		// encapsulate the fly method through airport.accept
		airport.Accept(flyingMachine)
	}
}
