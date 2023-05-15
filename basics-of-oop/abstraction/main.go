package main

import (
	"github.com/anarchymonkey/design-patterns/basics-of-oop/abstraction/book"
	driveAirplane "github.com/anarchymonkey/design-patterns/basics-of-oop/abstraction/drive-airplane"
)

func main() {
	// airplane booking application
	var airplane book.Airplane = book.Airplane{
		NoOfSeats:      41,
		SeatsAvailable: 1,
	}

	airplane.BookFlight(400)

	var flyingAirplane driveAirplane.Airplane = driveAirplane.Airplane{
		WingSpeed: 100,
		Direction: "WEST",
	}

	flyingAirplane.Fly(flyingAirplane.Direction)
}
