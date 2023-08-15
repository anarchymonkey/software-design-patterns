package director

import "github.com/anarchymonkey/design-patterns/creational/builder/builders"

type Director struct {
	builder builders.Builder
}

type SportsCar struct {
	ModelYear string
}

func (d *Director) SetBuilder(builder builders.Builder) {
	d.builder = builder
}

func (d *Director) GetSportsCar(builder builders.Builder) {
	builder.SetColor("White")
	builder.SetSeats(2)
	builder.SetGPS(55, 102)
	builder.SetEngine(builders.Engine{
		Horsepower: 1200,
		Title:      "IVTech Horse engine",
	})

}
