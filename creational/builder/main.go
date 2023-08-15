package main

import (
	"fmt"

	"github.com/anarchymonkey/design-patterns/creational/builder/builders"
	"github.com/anarchymonkey/design-patterns/creational/builder/director"
)

func main() {

	carBuilder := builders.ConcreteCarBuilder{}

	director := director.Director{}

	director.SetBuilder(&carBuilder)

	director.GetSportsCar(&carBuilder)

	fmt.Println(carBuilder.GetResult())
}
