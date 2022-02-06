package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kvloginov/t3oe/internal/base"
	"github.com/kvloginov/t3oe/internal/drawing"
)

type Platform struct {
	base.VolumeObject
	Team
}

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

	pixels := drawingStuff.ToPixels(p.Positional)
	io.GeoM.Translate(pixels.X, pixels.Y)
	io.GeoM.Scale(0.2, 0.2)

	screen.DrawImage(drawing.ROCKET_IMG, io)

}
