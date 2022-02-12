package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kvloginov/t3oe/internal/base"
	"github.com/kvloginov/t3oe/internal/drawing"
	"image/color"
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
			DragCoefficient:              0.999,
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
	io := &ebiten.DrawImageOptions{}

	//scale
	originalWidth, originalHeight := drawing.ROCKET_IMG.Size()
	requiredWidth := p.Width * float64(drawingStuff.UnitSize)
	requiredHeight := p.Height * float64(drawingStuff.UnitSize)
	scaleX := requiredWidth / float64(originalWidth)
	scaleY := requiredHeight / float64(originalHeight)
	io.GeoM.Scale(scaleX, scaleY)

	//positionalByPivot
	shiftForPivotX := -p.PivotRelativeX * p.Width
	shiftForPivotY := -p.PivotRelativeY * p.Height
	io.GeoM.Translate(drawingStuff.ToPixelsXY(shiftForPivotX, shiftForPivotY))

	//rotate
	io.GeoM.Rotate(p.Angle)

	//move to position
	io.GeoM.Translate(drawingStuff.ToPixelsXY(p.X, p.Y))

	screen.DrawImage(drawing.ROCKET_IMG, io)

	// position point
	pxls := drawingStuff.ToPixels(p.Positional)
	ebitenutil.DrawRect(screen, pxls.X, pxls.Y, 2, 2, color.RGBA{
		R: 0x00,
		G: 0xff,
		B: 0x00,
		A: 0xff,
	})
}

func (p *Platform) Update(dt float64) {
	// и должна быть скорость не 1
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
