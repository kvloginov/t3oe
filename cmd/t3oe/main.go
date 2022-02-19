// Copyright 2021 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
