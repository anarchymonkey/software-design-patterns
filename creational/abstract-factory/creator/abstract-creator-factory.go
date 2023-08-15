package creator

import (
	"github.com/anarchymonkey/design-patterns/creational/abstract-factory/product"
)

type FurnitureSelector interface {
	getChair() product.Chair
	getSofa() product.Sofa
	getTable() product.Table
}
