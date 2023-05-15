package book

import "fmt"

type Airplane struct {
	NoOfSeats      int
	SeatsAvailable int
}

func (airplane *Airplane) BookFlight(seatNo int) {
	if seatNo > airplane.NoOfSeats {
		fmt.Println("Seat number exceeds the no of seats")
		return
	}

	fmt.Println("Booked successfully")
}
