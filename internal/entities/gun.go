package entities

import (
	"github.com/kvloginov/t3oe/internal/base"
	"log"
	"time"
)

type Gun struct {
	canShoot   bool
	shootDelay time.Duration
	team       Team
}

func NewGun(shootDelay time.Duration, team Team) *Gun {
	return &Gun{true, shootDelay, team}
}

func (g *Gun) PullTheTrigger(shootFrom base.Positional) {
	if g.canShoot {
		log.Println("SHOT!")
		NewBullet(base.CopyPositional(shootFrom), g.team)

		g.canShoot = false
		go g.reload()
	}
}
func (g *Gun) reload() {
	timer := time.NewTimer(g.shootDelay)

	// TODO: context done
	select {
	case <-timer.C:
		g.canShoot = true
	}

}
