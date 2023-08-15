package main

import (
	"fmt"

	concreteproduct "github.com/anarchymonkey/design-patterns/creational/abstract-factory/concrete-product"
	"github.com/anarchymonkey/design-patterns/creational/abstract-factory/creator"
)

func main() {
	fmt.Println("This is the abstract factory implementation")

	for i := 0; i < 10; i = i + 2 {
		victorianFactory := creator.NewVictorianChairFactory{
			VictorianChair: concreteproduct.VictorianChair{
				NoOfLegs: i,
			},
			VictorianSofa: concreteproduct.VictorianSofa{
				NoOfLegs: i + 2,
			},
			VictorianTable: concreteproduct.VictorianTable{
				NoOfLegs: i + 4,
			},
		}

		chair := victorianFactory.GetChair()
		sofa := victorianFactory.GetSofa()
		table := victorianFactory.GetTable()

		const PREFIX = "desc: %s with no of legs %d \n"

		fmt.Printf(PREFIX, chair.GetDescription(), chair.GetNumberOfLegs())
		fmt.Printf(PREFIX, sofa.GetDescription(), sofa.GetNumberOfLegs())
		fmt.Printf(PREFIX, table.GetDescription(), table.GetNumberOfLegs())

		fmt.Println()
	}
}
