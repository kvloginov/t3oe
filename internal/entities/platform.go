package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kvloginov/t3oe/internal/base"
	"github.com/kvloginov/t3oe/internal/drawing"
	"image/color"
)

type Platform struct {
	base.VolumeObject
	Team
}

const FLY_MAX_SPEED = 5

func NewPlatform(positional base.Positional, team Team) Platform {
	return Platform{
		VolumeObject: base.VolumeObject{
			Positional:     positional,
			PivotRelativeX: 0.5,
			PivotRelativeY: 0.5,
			Width:          1,
			Height:         1,
		},
		Team: team}
}

func (p *Platform) Draw(screen *ebiten.Image, drawingStuff *drawing.DrawingStuff) {
	io := &ebiten.DrawImageOptions{}

	//TODO: Конечно, это нужно вынести в отдельную функцию в drawingStuff
	shiftX := p.PivotRelativeX * p.Width
	shiftY := p.PivotRelativeY * p.Height
	drawingStart := base.Positional{
		X: p.X - shiftX,
		Y: p.Y - shiftY,
	}
	pixels := drawingStuff.ToPixels(drawingStart)

	originalWidth, originalHeight := drawing.ROCKET_IMG.Size()
	requiredWidth := p.Width * float64(drawingStuff.UnitSize)
	requiredHeight := p.Height * float64(drawingStuff.UnitSize)

	scaleX := requiredWidth / float64(originalWidth)
	scaleY := requiredHeight / float64(originalHeight)

	io.GeoM.Scale(scaleX, scaleY)
	io.GeoM.Translate(pixels.X, pixels.Y)

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

	// TODO: Это, конечно, должно делаться через commander-а (класс, отдающий текущие действия, например), а не напрямую

	//TODO: В параметры тела
	p.SpeedX *= 0.95
	p.SpeedY *= 0.95

	// и должна быть скорость не 1
	if inpututil.KeyPressDuration(ebiten.KeyArrowUp) > 0 {
		p.SpeedY = -FLY_MAX_SPEED
	}
	if inpututil.KeyPressDuration(ebiten.KeyArrowDown) > 0 {
		p.SpeedY = FLY_MAX_SPEED
	}
	if inpututil.KeyPressDuration(ebiten.KeyArrowLeft) > 0 {
		p.SpeedX = -FLY_MAX_SPEED
	}
	if inpututil.KeyPressDuration(ebiten.KeyArrowRight) > 0 {
		p.SpeedX = FLY_MAX_SPEED
	}

	p.X += p.SpeedX * dt
	p.Y += p.SpeedY * dt

}
