package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kvloginov/t3oe/internal/base"
	"github.com/kvloginov/t3oe/internal/drawing"
	"math"
)

type Platform struct {
	base.Physical
	base.VolumeObject
	Team
}

const FLY_MAX_SPEED = 5

func NewPlatform(positional base.Positional, team Team) Platform {
	return Platform{
		Physical: base.Physical{
			Positional:                   positional,
			DragCoefficient:              0.965,
			TurningResistanceCoefficient: 0.950,
		},
		VolumeObject: base.VolumeObject{
			PivotRelativeX: 0.5,
			PivotRelativeY: 0.5,
			Width:          2,
			Height:         1,
		},
		Team: team}
}

func (p *Platform) Draw(screen *ebiten.Image, drawingStuff *drawing.DrawingStuff) {
	drawingStuff.DrawVolumeObject(p.VolumeObject, p.Positional, drawing.ROCKET_IMG, screen)
	drawingStuff.DrawDebugPositionPoint(p.Positional, screen)
}

func (p *Platform) Update(dt float64) {
	if inpututil.KeyPressDuration(ebiten.KeyArrowUp) > 0 {
		p.Speed = FLY_MAX_SPEED
	}
	if inpututil.KeyPressDuration(ebiten.KeyArrowDown) > 0 {
		p.Speed = -FLY_MAX_SPEED
	}
	if inpututil.KeyPressDuration(ebiten.KeyArrowLeft) > 0 {
		p.TurningSpeed = -math.Pi
	}
	if inpututil.KeyPressDuration(ebiten.KeyArrowRight) > 0 {
		p.TurningSpeed = math.Pi
	}

	p.Physical.Update(dt)
}
