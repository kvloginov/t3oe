package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kvloginov/t3oe/internal/base"
	"github.com/kvloginov/t3oe/internal/drawing"
	"math"
	"time"
)

type Platform struct {
	base.Physical
	base.VolumeObject
	Team

	controller  PlatformController
	rocketImage *ebiten.Image
	gun         *Gun
}

const FLY_MAX_SPEED = 5

func NewPlatform(positional base.Positional, team Team, controller PlatformController) *Platform {
	var rocketImage *ebiten.Image
	switch team {
	case TEAM_BLUE:
		rocketImage = drawing.ROCKET_BLUE_IMG
	case TEAM_RED:
		rocketImage = drawing.ROCKET_RED_IMG
	}

	return &Platform{
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
		Team:        team,
		controller:  controller,
		rocketImage: rocketImage,
		gun:         NewGun(time.Second/2, team),
	}
}

func (p *Platform) Draw(screen *ebiten.Image, drawingStuff *drawing.DrawingStuff) {
	drawingStuff.DrawVolumeObject(p.VolumeObject, p.Positional, p.rocketImage, screen)
	drawingStuff.DrawDebugPositionPoint(p.Positional, screen)
}

func (p *Platform) Update(dt float64) {
	if p.controller.Left() {
		p.TurningSpeed = -math.Pi
	} else if p.controller.Right() {
		p.TurningSpeed = math.Pi
	}

	if p.controller.Forward() {
		p.Speed = FLY_MAX_SPEED
	} else if p.controller.Backward() {
		p.Speed = -FLY_MAX_SPEED
	}

	if p.controller.Shoot() {
		p.gun.PullTheTrigger(p.Positional)
	}

	p.Physical.Update(dt)
}
