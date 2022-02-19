package drawing

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kvloginov/t3oe/internal/base"
	"github.com/kvloginov/t3oe/resources"
	"image/color"
)

type DrawingStuff struct {
	ScreenWidth      int
	ScreenHeight     int
	FieldUnitsWidth  int
	FieldUnitsHeight int
	UnitSize         int
}

var (
	ROCKET_BLUE_IMG *ebiten.Image
	ROCKET_RED_IMG  *ebiten.Image
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
	ROCKET_BLUE_IMG = getImageByPath("images/rocketRed.png")
	ROCKET_RED_IMG = getImageByPath("images/rocketBlue.png")
}

func getImageByPath(path string) *ebiten.Image {
	img, err := resources.Images.Open(path)
	if err != nil {
		panic(err)
	}

	ebitenImage, _, err := ebitenutil.NewImageFromReader(img)
	if err != nil {
		panic(err)
	}
	return ebitenImage
}

func (s *DrawingStuff) ToPixels(positional base.Positional) base.Positional {
	return base.NewPositional(positional.X*float64(s.UnitSize), positional.Y*float64(s.UnitSize), 0)
}

func (s *DrawingStuff) ToPixelsXY(x float64, y float64) (float64, float64) {
	return x * float64(s.UnitSize), y * float64(s.UnitSize)
}

func (s *DrawingStuff) DrawDebugPositionPoint(
	position base.Positional,
	screen *ebiten.Image,
) {
	// position point
	pxls := s.ToPixels(position)
	ebitenutil.DrawRect(screen, pxls.X, pxls.Y, 2, 2, color.RGBA{
		R: 0x00,
		G: 0xff,
		B: 0x00,
		A: 0xff,
	})
}

func (s *DrawingStuff) DrawVolumeObject(
	volumeObject base.VolumeObject,
	position base.Positional,
	rawImage *ebiten.Image,
	screen *ebiten.Image,
) {
	io := &ebiten.DrawImageOptions{}

	//scale
	originalWidth, originalHeight := rawImage.Size()
	requiredWidth := volumeObject.Width * float64(s.UnitSize)
	requiredHeight := volumeObject.Height * float64(s.UnitSize)
	scaleX := requiredWidth / float64(originalWidth)
	scaleY := requiredHeight / float64(originalHeight)
	io.GeoM.Scale(scaleX, scaleY)

	//positionalByPivot
	shiftForPivotX := -volumeObject.PivotRelativeX * volumeObject.Width
	shiftForPivotY := -volumeObject.PivotRelativeY * volumeObject.Height
	io.GeoM.Translate(s.ToPixelsXY(shiftForPivotX, shiftForPivotY))

	//rotate
	io.GeoM.Rotate(position.Angle)

	//move to position
	io.GeoM.Translate(s.ToPixelsXY(position.X, position.Y))

	screen.DrawImage(rawImage, io)
}
