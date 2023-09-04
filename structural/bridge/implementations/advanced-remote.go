package implementations

import "github.com/anarchymonkey/design-patterns/structural/bridge/abstractions"

type AdvancedRemote struct {
	Device abstractions.Device
}

func (advRemote *AdvancedRemote) TurnVolumeUp() {
	currentVolumeLevel := advRemote.Device.GetCurrentVolumeLevel() + 1
	advRemote.Device.SetVolumeLevel(currentVolumeLevel)
}

func (advRemote *AdvancedRemote) TurnVolumeDown() {
	currentVolumeLevel := advRemote.Device.GetCurrentVolumeLevel() - 1
	advRemote.Device.SetVolumeLevel(currentVolumeLevel)
}

func (advRemote *AdvancedRemote) GoChannelUp() {
	advRemote.Device.SetCurrentChannelNumber(advRemote.Device.GetCurrentChannel() + 1)
}

func (advRemote *AdvancedRemote) GoChannelDown() {
	advRemote.Device.SetCurrentChannelNumber(advRemote.Device.GetCurrentChannel() - 1)
}

func (advRemote *AdvancedRemote) Start() {
	advRemote.Device.TurnOn()
}

func (advRemote *AdvancedRemote) Stop() {
	advRemote.Device.TurnOff()
}

func (advRemote *AdvancedRemote) Mute() {
	advRemote.Device.SetVolumeLevel(0)
}
