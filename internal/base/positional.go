package base

import "math"

const (
	ANGLE_LEFT  = math.Pi
	ANGLE_RIGHT = 0
	ANGLE_DOWN  = math.Pi / 2
	ANGLE_UP    = (math.Pi * 3) / 2
)

type Positional struct {
	X     float64 // units
	Y     float64 // units
	Angle float64 // In radians
}

func NewPositional(x float64, y float64, angle float64) Positional {
	return Positional{X: x, Y: y, Angle: angle}
}

func CopyPositional(from Positional) Positional {
	// ok we could just return `from`
	return Positional{X: from.X, Y: from.Y, Angle: from.Angle}
}
