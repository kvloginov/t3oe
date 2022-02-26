package gameObjects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kvloginov/t3oe/internal/drawing"
)

type Drawable interface {
	Draw(screen *ebiten.Image, drawingStuff *drawing.DrawingStuff)
}
