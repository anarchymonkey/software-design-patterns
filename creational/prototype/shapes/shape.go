package shapes

const DEFAULT_TYPE = "_original"

type Shape interface {
	Init(x, y int, color string)
	Clone() Shape
}
