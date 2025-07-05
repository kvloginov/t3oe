package entities

import (
	"math"
	"time"

	"github.com/kvloginov/t3oe/internal/base"
)

type Thruster struct {
	team              Team
	lastParticleTime  time.Time
	particleInterval  time.Duration
	particlesPerBurst int
}

func NewThruster(team Team) *Thruster {
	return &Thruster{
		team:              team,
		lastParticleTime:  time.Time{},
		particleInterval:  time.Millisecond * 25,
		particlesPerBurst: 3,
	}
}

func (t *Thruster) Activate(thrusterPos base.Positional) {
	now := time.Now()
	if now.Sub(t.lastParticleTime) >= t.particleInterval {
		t.lastParticleTime = now

		// Create several particles at once
		for i := 0; i < t.particlesPerBurst; i++ {
			// Particles should appear behind the platform
			// Shift position backward relative to platform direction
			backwardOffset := 0.5 // Slightly overlaps with platform
			particlePos := base.Positional{
				Pos: thrusterPos.Pos.Plus(
					base.NewVectorWithAngle(thrusterPos.Angle + math.Pi).MultiplyScalar(backwardOffset),
				),
				Angle: thrusterPos.Angle + math.Pi, // Particles fly backward
			}

			NewThrusterParticle(particlePos)
		}
	}
}
