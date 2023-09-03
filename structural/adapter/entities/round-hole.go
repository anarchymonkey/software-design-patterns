package entities

func (roundHole *RoundHole) GetRadius() float32 {
	return roundHole.Radius
}

func (roundHole *RoundHole) Fits(peg RoundPeg) bool {
	return roundHole.Radius >= peg.GetRadius()
}
