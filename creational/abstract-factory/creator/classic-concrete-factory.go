package creator

import concreteproduct "github.com/anarchymonkey/design-patterns/creational/abstract-factory/concrete-product"

type ClassicCreator interface {
	GetChair() concreteproduct.ClassicChair
	GetSofa() concreteproduct.ClassicSofa
	GetTable() concreteproduct.ClassicTable
}
