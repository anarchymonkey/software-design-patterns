package creator

import concreteproduct "github.com/anarchymonkey/design-patterns/creational/abstract-factory/concrete-product"

type MordernFactory interface {
	GetChair() concreteproduct.MordernChair
	GetSofa() concreteproduct.MordernSofa
	GetTable() concreteproduct.MordernTable
}
