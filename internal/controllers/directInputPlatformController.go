package controllers

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type directInputPlatformController struct {
}

func NewDirectInputPlatformController() *directInputPlatformController {
	return &directInputPlatformController{}
}

func (d *directInputPlatformController) Forward() bool {
	return inpututil.KeyPressDuration(ebiten.KeyArrowUp) > 0
}

func (d *directInputPlatformController) Backward() bool {
	return inpututil.KeyPressDuration(ebiten.KeyArrowDown) > 0
}

func (d *directInputPlatformController) Left() bool {
	return inpututil.KeyPressDuration(ebiten.KeyArrowLeft) > 0
}

func (d *directInputPlatformController) Right() bool {
	return inpututil.KeyPressDuration(ebiten.KeyArrowRight) > 0
}
