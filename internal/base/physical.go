package base

type Physical struct {
	Positional

	Speed           Vector
	Acceleration    Vector
	DragCoefficient float64

	TurningSpeed                 float64 // radians/sec
	Torque                       float64 // radians/sec2
	TurningResistanceCoefficient float64
}

func (p *Physical) Update(dt float64) {
	// moving
	// A
	if p.DragCoefficient != 0 {
		spLen := p.Speed.Length()
		force := p.DragCoefficient * spLen * spLen
		acc := p.Speed.WithLength(force)
		p.Acceleration = p.Acceleration.Minus(acc)
	}
	// V
	p.Speed = p.Speed.Plus(p.Acceleration.MultiplyScalar(dt))
	// X
	p.Pos = p.Pos.Plus(p.Speed.MultiplyScalar(dt))

	// turning
	// ε
	// ω
	if p.TurningResistanceCoefficient != 0 {
		p.TurningSpeed *= p.TurningResistanceCoefficient
	}
	p.TurningSpeed += p.Torque * dt
	// φ
	p.Positional.Angle += p.TurningSpeed * dt

	// reset acceleration
	p.Acceleration = Vector{
		X: 0,
		Y: 0,
	}
}

func NewPhysicalByPos(positional Positional) *Physical {
	return &Physical{Positional: positional}
}
