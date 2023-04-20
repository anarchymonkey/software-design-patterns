package main

import (
	"github.com/anarchymonkey/design-patterns/creational/factory/creator"
)

func main() {
	var transport creator.Transport

	transport = creator.TransportFactory("Truck", 200)
	transport.Deliver()

	transport = creator.TransportFactory("Ship", 20)
	transport.Deliver()
}
