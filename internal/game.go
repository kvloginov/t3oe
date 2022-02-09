package internal

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
		platforms: make([]entities.Platform, 1, 1),
		bullets:   make([]entities.Bullet, 0),
		drawingStuff: drawing.NewDrawingStuff(
			screenWidth,
			screenHeight,
			fieldUnitsWidth,
			fieldUnitsHeight,
			unitSize),
	}

	g.platforms[0] = entities.NewPlatform(base.NewMovable(float64(fieldUnitsWidth/2), float64(fieldUnitsHeight-1)), entities.TEAM_DOWN)
	//g.platforms[1] = entities.NewPlatform(base.NewMovable(12, 5), entities.TEAM_DOWN)

	return g
}

func (g *Game) Update() error {
	//TODO: не использовать эту функцию т.к. написано, что только для дебуга.
	tps := ebiten.CurrentTPS()
	dt := 1 / tps

	for i, _ := range g.platforms {
		g.platforms[i].Update(dt)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprint(ebiten.CurrentTPS()))
	for i, _ := range g.platforms {
		g.platforms[i].Draw(screen, g.drawingStuff)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.drawingStuff.ScreenWidth, g.drawingStuff.ScreenHeight
}
