package abstractions

type Device interface {
	TurnOn()
	TurnOff()
	GetCurrentVolumeLevel() int
	SetVolumeLevel(volumeLevel int)
	GetCurrentChannel() int
	SetCurrentChannelNumber(channelNumber int)
}
