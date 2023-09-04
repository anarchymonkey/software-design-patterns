package abstractions

type Remote interface {
	TurnVolumeUp()
	TurnVolumeDown()
	GoChannelUp()
	GoChannelDown()
	Start()
	Stop()
}
