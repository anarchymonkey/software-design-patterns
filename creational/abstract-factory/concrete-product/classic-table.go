package concreteproduct

type ClassicTable struct {
	NoOfLegs int
}

func (c *ClassicTable) GetNumberOfLegs() int {
	return c.NoOfLegs
}

func (c *ClassicTable) GetDescription() string {
	return "This is a mordern chair"
}
