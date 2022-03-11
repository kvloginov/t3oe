package controllers

import (
	"math/rand"
	"time"
)

type turning int

const (
	TURNING_NOTHING turning = iota
	TURNING_LEFT
	TURNING_RIGHT
)

type movement int

const (
	MOVEMENT_NOTHING movement = iota
	MOVEMENT_FORWARD
	MOVEMENT_BACKWARD
)

type randomPlatformController struct {
	currentMovement movement
	currentTurning  turning
}

func (d *randomPlatformController) Shoot() bool {
	return false
}

func NewRandomPlatformController() *randomPlatformController {
	controller := &randomPlatformController{}

	forwardTicker := time.NewTicker(time.Second)
	turningTicker := time.NewTicker(time.Second / 2)

	// TODO: add graceful shutdown
	go controller.handleTicks(forwardTicker, turningTicker)

	return controller
}

func (d *randomPlatformController) handleTicks(forwardTicker *time.Ticker, turningTicker *time.Ticker) {
	for {
		select {
		case <-forwardTicker.C:
			d.currentMovement = randomMovement()
		case <-turningTicker.C:
			d.currentTurning = randomTurning()
		}
	}
}

func randomMovement() movement {
	rnd := rand.Intn(100)
	if betweenInt(rnd, 0, 60) {
		return MOVEMENT_FORWARD
	}
	if betweenInt(rnd, 60, 80) {
		return MOVEMENT_BACKWARD
	}
	return MOVEMENT_NOTHING
}

func randomTurning() turning {
	rnd := rand.Intn(100)
	if betweenInt(rnd, 0, 60) {
		return TURNING_NOTHING
	}
	if betweenInt(rnd, 60, 80) {
		return TURNING_LEFT
	}
	return TURNING_RIGHT
}

// is value in [first, second)
func betweenInt(value, first, second int) bool {
	return value >= first && value < second
}

func (d *randomPlatformController) Forward() bool {
	return d.currentMovement == MOVEMENT_FORWARD
}

func (d *randomPlatformController) Backward() bool {
	return d.currentMovement == MOVEMENT_BACKWARD
}

func (d *randomPlatformController) Left() bool {
	return d.currentTurning == TURNING_LEFT
}

func (d *randomPlatformController) Right() bool {
	return d.currentTurning == TURNING_RIGHT
}
