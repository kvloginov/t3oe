package entities

import (
	"github.com/kvloginov/t3oe/internal/base"
)

type Bullet struct {
	base.Positional
	Team
}

func NewBullet(movable base.Positional, team Team) *Bullet {
	return &Bullet{Positional: movable, Team: team}
}
