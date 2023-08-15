package concreteproduct

type ClassicSofa struct {
	NoOfLegs int
}

func (c *ClassicSofa) GetNumberOfLegs() int {
	return c.NoOfLegs
}

func (c *ClassicSofa) GetDescription() string {
	return "This is a mordern chair"
}
