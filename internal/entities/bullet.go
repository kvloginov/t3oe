package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kvloginov/t3oe/internal/base"
	"github.com/kvloginov/t3oe/internal/drawing"
)

const SIMPLE_BULLET_SPEED = 20

type Bullet struct {
	base.VolumeObject
	base.Physical
	Team
}

func NewBullet(pos base.Positional, team Team) *Bullet {
	return &Bullet{
		Physical: base.Physical{
			Positional: pos,
			Speed:      SIMPLE_BULLET_SPEED,
		},
		VolumeObject: base.VolumeObject{
			PivotRelativeX: 0.5,
			PivotRelativeY: 0.5,
			Width:          24 / float64(64),
			Height:         16 / float64(64),
		},
		Team: team}
}

func (p *Bullet) Draw(screen *ebiten.Image, drawingStuff *drawing.DrawingStuff) {
	drawingStuff.DrawVolumeObject(p.VolumeObject, p.Positional, drawing.SIMPLE_BULLET_IMG, screen)
	drawingStuff.DrawDebugPositionPoint(p.Positional, screen)
}

func (p *Bullet) Update(dt float64) {
	p.Physical.Update(dt)
}
