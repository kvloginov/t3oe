package base

import "math"

type Physical struct {
	Positional

	Speed float64
	//SpeedX       float64 // units/sec
	//SpeedY       float64 // units/sec
	TurningSpeed float64 // radians/sec

	Acceleration float64
	//AccelerationX float64 // units/sec2
	//AccelerationY float64 // units/sec2
	Torque float64 // radians/sec2

	DragCoefficient              float64 // just * acceleration to X and Y
	TurningResistanceCoefficient float64
}

func (p *Physical) Update(dt float64) {
	//p.SpeedX += p.AccelerationX * dt
	//p.SpeedY += p.AccelerationY * dt

	p.Speed += p.Acceleration * dt

	if p.DragCoefficient != 0 {
		p.Speed *= p.DragCoefficient
		//p.SpeedX *= p.DragCoefficient
		//p.SpeedY *= p.DragCoefficient
	}

	p.Positional.X += p.Speed * math.Cos(p.Positional.Angle) * dt
	p.Positional.Y += p.Speed * math.Sin(p.Positional.Angle) * dt
	//p.Positional.X += p.SpeedX * dt
	//p.Positional.Y += p.SpeedY * dt

	if p.TurningResistanceCoefficient != 0 {
		p.TurningSpeed *= p.TurningResistanceCoefficient
	}
	p.TurningSpeed += p.Torque * dt
	p.Positional.Angle += p.TurningSpeed * dt
}

func NewPhysicalByPos(positional Positional) *Physical {
	return &Physical{Positional: positional}
}
