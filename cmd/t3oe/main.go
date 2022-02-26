package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kvloginov/t3oe/internal"
	"log"
	"math/rand"
	"time"
)

const (
	fieldUnitsWidth  = 20
	fieldUnitsHeight = 15

	unitSize = 64
)

func main() {
	rand.Seed(time.Now().UnixNano())
	width := fieldUnitsWidth * unitSize
	height := fieldUnitsHeight * unitSize

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("DUEL t3oe")
	if err := ebiten.RunGame(internal.NewGame(width, height, fieldUnitsWidth, fieldUnitsHeight, unitSize)); err != nil {
		log.Fatal(err)
	}
}
