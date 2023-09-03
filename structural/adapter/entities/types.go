package entities

type Hole interface {
	getRadius() float32
	fits() bool
}

type RoundPeg struct {
	Radius float32
}

type SquarePeg struct {
	Width float32
}

type RoundHole struct {
	Radius float32
}
