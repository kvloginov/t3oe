package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kvloginov/t3oe/internal/base"
	"github.com/kvloginov/t3oe/internal/drawing"
	"github.com/kvloginov/t3oe/internal/gameObjects"
)

type Explosion struct {
	base.VolumeObject
	base.Positional
	NamedEntity
}

func NewExplosion(pos base.Positional) *Explosion {
	exp := &Explosion{
		VolumeObject: base.VolumeObject{},
		Positional:   base.Positional{},
		NamedEntity:  NamedEntity{},
	}

	gameObjects.GameObjects.RegisterWithGeneratedId(exp)
	return exp
}

func (p *Explosion) Draw(screen *ebiten.Image, drawingStuff *drawing.DrawingStuff) {
	drawingStuff.DrawVolumeObject(p.VolumeObject, p.Positional, drawing.SIMPLE_BULLET_IMG, screen)
	drawingStuff.DrawDebugPositionPoint(p.Positional, screen)
}

func (p *Explosion) Update(dt float64) {
}
