package drawing

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kvloginov/t3oe/internal/base"
	"github.com/kvloginov/t3oe/resources"
)

type DrawingStuff struct {
	ScreenWidth      int
	ScreenHeight     int
	FieldUnitsWidth  int
	FieldUnitsHeight int
	UnitSize         int
}

var (
	ROCKET_IMG *ebiten.Image
)

func NewDrawingStuff(screenWidth int, screenHeight int, fieldUnitsWidth int, fieldUnitsHeight int, unitSize int) *DrawingStuff {
	drSt := &DrawingStuff{
		ScreenWidth:      screenWidth,
		ScreenHeight:     screenHeight,
		FieldUnitsWidth:  fieldUnitsWidth,
		FieldUnitsHeight: fieldUnitsHeight,
		UnitSize:         unitSize,
	}

	initImages()

	return drSt
}

func initImages() {
	file, err := resources.Images.Open("images/rocketUp.png")
	if err != nil {
		panic(err)
	}

	ROCKET_IMG, _, err = ebitenutil.NewImageFromReader(file)
	if err != nil {
		return
	}
}

func (s *DrawingStuff) ToPixels(positional base.Positional) base.Positional {
	return base.NewMovable(positional.X*float64(s.UnitSize), positional.Y*float64(s.UnitSize))
}

func (s *DrawingStuff) ToGeoM(volumeObj base.VolumeObject) ebiten.GeoM {
	geom := ebiten.GeoM{}
	geom.Rotate(volumeObj.Angle)

	//TODO: Продолжить тут
	return ebiten.GeoM{}
	//return base.NewMovable(movable.X*float64(s.UnitSize), movable.Y*float64(s.UnitSize))
}
