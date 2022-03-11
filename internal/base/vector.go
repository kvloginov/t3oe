package base

import (
	"fmt"
	"math"
)

type Vector struct {
	X float64
	Y float64
}

func (v Vector) Plus(v2 Vector) Vector {
	return Vector{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v Vector) Minus(v2 Vector) Vector {
	return Vector{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

func NewVectorWithAngle(angle float64) Vector {
	vector := &Vector{
		X: math.Cos(angle),
		Y: math.Sin(angle),
	}
	return vector.Unit()
}

func (v Vector) Unit() Vector {
	length := v.Length()
	if length == 0 {
		return Vector{
			X: 1,
			Y: 0,
		}
	}
	return Vector{
		X: v.X / length,
		Y: v.Y / length,
	}
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector) WithLength(len float64) Vector {
	unit := v.Unit()
	return Vector{
		X: unit.X * len,
		Y: unit.Y * len,
	}
}

func (v Vector) MultiplyScalar(multiplier float64) Vector {
	return Vector{
		X: v.X * multiplier,
		Y: v.Y * multiplier,
	}
}

func (v Vector) String() string {
	return fmt.Sprintf("X: %v, Y: %v", v.X, v.Y)
}
