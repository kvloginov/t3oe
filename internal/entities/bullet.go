package entities

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kvloginov/t3oe/internal/base"
	"github.com/kvloginov/t3oe/internal/drawing"
	"github.com/kvloginov/t3oe/internal/gameObjects"
	"github.com/kvloginov/t3oe/internal/trigger"
)

const SIMPLE_BULLET_SPEED = 20

type Bullet struct {
	base.VolumeObject
	base.Physical
	Team
	NamedEntity
}

func NewBullet(pos base.Positional, team Team) *Bullet {
	bul := &Bullet{
		Physical: base.Physical{
			Positional: pos,
			Speed:      base.NewVectorWithAngle(pos.Angle).MultiplyScalar(SIMPLE_BULLET_SPEED),
		},
		VolumeObject: base.VolumeObject{
			PivotRelativeX: 0.5,
			PivotRelativeY: 0.5,
			Width:          24 / float64(64),
			Height:         16 / float64(64),
		},
		Team: team}

	id := gameObjects.GameObjects.RegisterWithGeneratedId(bul)
	trigger.RegisterTrigger(bul, bul.Width, func(object1 interface{}, object2 interface{}) {
		switch object2.(type) {
		case HasTeam:
			if object2.(HasTeam).GetTeam() != bul.GetTeam() {
				fmt.Printf("BABAH!")
				gameObjects.GameObjects.Destroy(id)
				trigger.DeleteTriggers(id)
			}
		}
	})

	return bul
}

func (p *Bullet) Draw(screen *ebiten.Image, drawingStuff *drawing.DrawingStuff) {
	drawingStuff.DrawVolumeObject(p.VolumeObject, p.Positional, drawing.SIMPLE_BULLET_IMG, screen)
	drawingStuff.DrawDebugPositionPoint(p.Positional, screen)
}

func (p *Bullet) Update(dt float64) {
	p.Physical.Update(dt)
}
