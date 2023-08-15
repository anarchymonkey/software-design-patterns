package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/anarchymonkey/design-patterns/creational/factory/creator"
)

/*

Factory pattern has 4 main components:

* Product
* Concrete Product
* Creator
* Concrete Creator

*/

func readLine() (string, string, string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter your from destination")
	scanner.Scan()
	from := scanner.Text()

	fmt.Println("Where are you going to")
	scanner.Scan()
	to := scanner.Text()

	fmt.Println("What is the message that you are going to deliver")
	scanner.Scan()
	message := scanner.Text()

	return from, to, message
}

func main() {

	from, to, message := readLine()

	const RoadTransportType = creator.TRUCK
	const SeaTransportType = creator.SHIP

	roadResp, err := creator.RoadTransportFactory(RoadTransportType)

	if err != nil {
		os.Exit(0)
		return
	}

	seaResp, err := creator.SeaTransportFactory(SeaTransportType)

	fmt.Println(roadResp.Deliver(from, to, message))
	fmt.Println(seaResp.Deliver(from, to, message))
}
