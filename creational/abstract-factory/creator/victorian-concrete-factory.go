package creator

import (
	concreteProduct "github.com/anarchymonkey/design-patterns/creational/abstract-factory/concrete-product"
	"github.com/anarchymonkey/design-patterns/creational/abstract-factory/product"
)

type NewVictorianChairFactory struct {
	VictorianChair concreteProduct.VictorianChair
	VictorianSofa  concreteProduct.VictorianSofa
	VictorianTable concreteProduct.VictorianTable
}

func (vf *NewVictorianChairFactory) GetChair() product.Chair {
	return &vf.VictorianChair
}

func (vf *NewVictorianChairFactory) GetSofa() product.Sofa {
	return &vf.VictorianSofa
}

func (vf *NewVictorianChairFactory) GetTable() product.Table {
	return &vf.VictorianTable
}
