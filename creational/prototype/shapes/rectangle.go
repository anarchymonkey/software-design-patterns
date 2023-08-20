package shapes

type Rectangle struct {
	Width  int
	Height int
	Color  string
	Type   string
}

func (r *Rectangle) Init(width, height int, color string) {
	r.Width = width
	r.Height = height
	r.Color = color
	r.Type = DEFAULT_TYPE

}

func (r *Rectangle) Clone() Shape {
	return &Rectangle{
		Width:  r.Width,
		Height: r.Height,
		Color:  r.Color,
		Type:   "_clone",
	}
}
