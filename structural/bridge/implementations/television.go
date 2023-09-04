package implementations

import "fmt"

type TV struct {
	IsOn               bool
	CurrentChannel     int
	CurrentVolumeLevel int
}

func (tv *TV) TurnOn() {
	fmt.Println("Turning on Tv")
	tv.IsOn = true
}

func (tv *TV) TurnOff() {
	fmt.Println("Turning off Tv")
	tv.IsOn = false
}

func (tv *TV) GetCurrentVolumeLevel() int {
	return tv.CurrentVolumeLevel
}

func (tv *TV) SetVolumeLevel(volumeLevel int) {
	tv.CurrentVolumeLevel = volumeLevel
}

func (tv *TV) GetCurrentChannel() int {
	return tv.CurrentChannel
}

func (tv *TV) SetCurrentChannelNumber(channelNumber int) {
	tv.CurrentChannel = channelNumber
}
