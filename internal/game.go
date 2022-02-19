package internal

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kvloginov/t3oe/internal/base"
	"github.com/kvloginov/t3oe/internal/controllers"
	"github.com/kvloginov/t3oe/internal/drawing"
	"github.com/kvloginov/t3oe/internal/entities"
)

type Game struct {
	gameObjects []GameObject

	drawingStuff *drawing.DrawingStuff
}

func NewGame(screenWidth int, screenHeight int, fieldUnitsWidth int, fieldUnitsHeight int, unitSize int) *Game {
	g := &Game{
		drawingStuff: drawing.NewDrawingStuff(
			screenWidth,
			screenHeight,
			fieldUnitsWidth,
			fieldUnitsHeight,
			unitSize),
	}

	g.gameObjects = append(g.gameObjects, entities.NewPlatform(
		base.NewPositional(float64(fieldUnitsWidth/2), float64(fieldUnitsHeight-1), base.ANGLE_UP),
		entities.TEAM_BLUE,
		controllers.NewDirectInputPlatformController(),
	))

	g.gameObjects = append(g.gameObjects, entities.NewPlatform(
		base.NewPositional(float64(fieldUnitsWidth/2), float64(fieldUnitsHeight/2), base.ANGLE_DOWN),
		entities.TEAM_RED,
		controllers.NewRandomPlatformController(),
	))

	return g
}

func (g *Game) Update() error {
	//TODO: не всегда 60 фпс
	dt := float64(1) / 60

	for i, _ := range g.gameObjects {
		g.gameObjects[i].Update(dt)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprint(ebiten.CurrentTPS()))

	for i, _ := range g.gameObjects {
		g.gameObjects[i].Draw(screen, g.drawingStuff)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.drawingStuff.ScreenWidth, g.drawingStuff.ScreenHeight
}
