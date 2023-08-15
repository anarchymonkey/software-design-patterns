package builders

import "fmt"

type CarModel struct {
	Color         string
	Seats         int32
	ModelYear     int32
	Name          string
	EngineDetails Engine
}

type ConcreteCarBuilder struct {
	Car CarModel
}

func (cCarBuilder *ConcreteCarBuilder) SetColor(color string) {
	cCarBuilder.Car.Color = color
}

func (cCarBuilder *ConcreteCarBuilder) SetEngine(engine Engine) {
	cCarBuilder.Car.EngineDetails = engine
}

func (cCarBuilder *ConcreteCarBuilder) SetGPS(latitude, longitude rune) {
	fmt.Println("The car has gps which starts from latitude & longitude", latitude, longitude)
}

func (cCarBuilder *ConcreteCarBuilder) SetSeats(noOfSeats int32) {
	cCarBuilder.Car.Seats = noOfSeats
}

func (cCarBuilder *ConcreteCarBuilder) GetResult() CarModel {
	return cCarBuilder.Car
}
