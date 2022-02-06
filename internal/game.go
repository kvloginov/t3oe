package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kvloginov/t3oe/internal/base"
	"github.com/kvloginov/t3oe/internal/drawing"
	"github.com/kvloginov/t3oe/internal/entities"
)

type Game struct {
	platforms []entities.Platform
	bullets   []entities.Bullet

	drawingStuff *drawing.DrawingStuff
}

func NewGame(screenWidth int, screenHeight int, fieldUnitsWidth int, fieldUnitsHeight int, unitSize int) *Game {
	g := &Game{
		platforms: make([]entities.Platform, 2, 2),
		bullets:   make([]entities.Bullet, 0),
		drawingStuff: drawing.NewDrawingStuff(
			screenWidth,
			screenHeight,
			fieldUnitsWidth,
			fieldUnitsHeight,
			unitSize),
	}

	g.platforms[0] = entities.NewPlatform(base.NewMovable(10, 1), entities.TEAM_UP)
	g.platforms[1] = entities.NewPlatform(base.NewMovable(12, 5), entities.TEAM_DOWN)

	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, platform := range g.platforms {
		platform.Draw(screen, g.drawingStuff)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.drawingStuff.ScreenWidth, g.drawingStuff.ScreenHeight
}
