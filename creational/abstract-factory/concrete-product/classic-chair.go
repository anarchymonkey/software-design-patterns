package concreteproduct

type ClassicChair struct {
	NoOfLegs int
}

func (c *ClassicChair) GetNumberOfLegs() int {
	return c.NoOfLegs
}

func (c *ClassicChair) GetDescription() string {
	return "This is a mordern chair"
}
