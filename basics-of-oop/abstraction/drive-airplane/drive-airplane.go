package driveAirplane

import "fmt"

type Airplane struct {
	WingSpeed int
	Direction string
}

func (airplane *Airplane) Fly(direction string) {
	fmt.Printf("The airplane is flying with a speed of %d in direction %s", airplane.WingSpeed, airplane.Direction)
}
