package shapes

type Circle struct {
	Radius float32
	Color  string
	Type   string
}

func (c *Circle) Init(x, y int, color string) {
	c.Color = color
	c.Radius = c.calculateRadius(x, y)
	c.Type = DEFAULT_TYPE
}

func (c *Circle) calculateRadius(x, y int) float32 {
	return float32(x + y)
}

func (c *Circle) Clone() Shape {
	return &Circle{
		Radius: c.Radius,
		Color:  c.Color,
		Type:   "_clone",
	}
}
