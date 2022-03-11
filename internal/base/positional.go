package base

import "math"

const (
	ANGLE_LEFT  = math.Pi
	ANGLE_RIGHT = 0
	ANGLE_DOWN  = math.Pi / 2
	ANGLE_UP    = (math.Pi * 3) / 2
)

type Positional struct {
	Pos   Vector
	Angle float64 // In radians
}

func NewPositional(x float64, y float64, angle float64) Positional {
	return Positional{
		Pos: Vector{
			X: x,
			Y: y,
		},
		Angle: angle}
}

func CopyPositional(from Positional) Positional {
	// ok we could just return `from`
	return Positional{
		Pos: Vector{
			X: from.Pos.X,
			Y: from.Pos.Y,
		},
		Angle: from.Angle,
	}
}

func (p *Positional) GetPosition() Vector {
	return p.Pos
}
