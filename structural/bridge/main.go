package main

import (
	"fmt"

	"github.com/anarchymonkey/design-patterns/structural/bridge/abstractions"
	"github.com/anarchymonkey/design-patterns/structural/bridge/implementations"
)

func main() {
	var device abstractions.Device

	tv := implementations.TV{
		IsOn:               true,
		CurrentChannel:     0,
		CurrentVolumeLevel: 50,
	}

	device = &tv

	device.GetCurrentChannel()

	var remote abstractions.Remote

	advancedRemote := implementations.AdvancedRemote{
		Device: device,
	}

	remote = &advancedRemote

	remote.Start()
	defer remote.Stop()

	remote.GoChannelUp()
	remote.GoChannelDown()
	remote.GoChannelDown()
	remote.GoChannelUp()
	remote.GoChannelUp()
	remote.GoChannelUp()
	remote.TurnVolumeDown()
	remote.TurnVolumeDown()
	remote.TurnVolumeUp()

	advancedRemote.Mute()

	fmt.Println("Current volume level is", device.GetCurrentVolumeLevel())
	fmt.Println("Current channel number is", device.GetCurrentChannel())

}
