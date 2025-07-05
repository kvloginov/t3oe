package entities

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kvloginov/t3oe/internal/base"
	"github.com/kvloginov/t3oe/internal/drawing"
	"github.com/kvloginov/t3oe/internal/gameObjects"
)

type ThrusterParticle struct {
	base.Physical
	base.VolumeObject
	NamedEntity

	lifeTime    float64
	maxLifeTime float64
	color       color.RGBA
	size        float64
}

func NewThrusterParticle(pos base.Positional) *ThrusterParticle {
	// Add random deviation to position and direction
	randomAngle := pos.Angle + (rand.Float64()-0.5)*math.Pi/4
	randomSpeed := 5.0 + rand.Float64()*2.0

	// Create random fire color (orange/red/yellow)
	fireColor := color.RGBA{
		R: uint8(200 + rand.Intn(55)), // 200-255
		G: uint8(50 + rand.Intn(100)), // 50-150
		B: uint8(rand.Intn(50)),       // 0-50
		A: 255,
	}

	// Random particle size
	particleSize := 0.2 + rand.Float64()*0.2

	particle := &ThrusterParticle{
		Physical: base.Physical{
			Positional:      pos,
			Speed:           base.NewVectorWithAngle(randomAngle).MultiplyScalar(randomSpeed),
			DragCoefficient: 0.8,
		},
		VolumeObject: base.VolumeObject{
			PivotRelativeX: 0.5,
			PivotRelativeY: 0.5,
			Width:          particleSize,
			Height:         particleSize,
		},
		maxLifeTime: 0.3 + rand.Float64()*0.3,
		lifeTime:    0,
		color:       fireColor,
		size:        particleSize,
	}

	particle.name = gameObjects.GameObjects.RegisterWithGeneratedId(particle)

	return particle
}

func (p *ThrusterParticle) Update(dt float64) {
	p.lifeTime += dt

	// Check lifetime and destroy if expired
	if p.lifeTime >= p.maxLifeTime {
		gameObjects.GameObjects.Destroy(p.GetName())
		return
	}

	// Update physics
	p.Physical.Update(dt)

	// Change color over time (becomes darker)
	lifeRatio := p.lifeTime / p.maxLifeTime
	if lifeRatio > 1.0 {
		lifeRatio = 1.0
	}

	// Gradually fade out
	alpha := 1.0 - lifeRatio
	p.color.A = uint8(255 * alpha)

	// Decrease size over time
	p.size = p.VolumeObject.Width * (1.0 - lifeRatio*0.5)
}

func (p *ThrusterParticle) Draw(screen *ebiten.Image, drawingStuff *drawing.DrawingStuff) {
	if p.color.A == 0 {
		return // Don't render fully transparent particles
	}

	// Draw particle as a rectangle
	pixelX, pixelY := drawingStuff.ToPixelsXY(p.Positional.Pos.X, p.Positional.Pos.Y)
	pixelSize := p.size * float64(drawingStuff.UnitSize)

	// Center the particle
	x := pixelX - pixelSize/2
	y := pixelY - pixelSize/2

	ebitenutil.DrawRect(screen, x, y, pixelSize, pixelSize, p.color)
}
