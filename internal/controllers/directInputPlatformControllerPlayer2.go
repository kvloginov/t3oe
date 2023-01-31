package controllers

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type directInputPlatformControllerPlayer2 struct {
}

func NewDirectInputPlatformControllerPlayer2() *directInputPlatformControllerPlayer2 {
	return &directInputPlatformControllerPlayer2{}
}

func (d *directInputPlatformControllerPlayer2) Forward() bool {
	return inpututil.KeyPressDuration(ebiten.KeyW) > 0
}

func (d *directInputPlatformControllerPlayer2) Backward() bool {
	return inpututil.KeyPressDuration(ebiten.KeyS) > 0
}

func (d *directInputPlatformControllerPlayer2) Left() bool {
	return inpututil.KeyPressDuration(ebiten.KeyA) > 0
}

func (d *directInputPlatformControllerPlayer2) Right() bool {
	return inpututil.KeyPressDuration(ebiten.KeyD) > 0
}

func (d *directInputPlatformControllerPlayer2) Shoot() bool {
	return inpututil.KeyPressDuration(ebiten.KeySpace) > 0
}
