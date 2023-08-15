package builders

type Engine struct {
	Horsepower float64
	Title      string
}
type Builder interface {
	SetColor(color string)
	SetSeats(noOfSeats int32)
	SetEngine(engine Engine)
	SetGPS(latitude, longitude rune)
}
