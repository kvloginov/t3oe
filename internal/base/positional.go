package base

type Positional struct {
	X      float64
	Y      float64
	SpeedX float64
	SpeedY float64
	Angle  float64 // In radians
}

func NewMovable(x float64, y float64) Positional {
	return Positional{X: x, Y: y, SpeedX: 0, SpeedY: 0, Angle: 0}
}
