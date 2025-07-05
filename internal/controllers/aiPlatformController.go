package controllers

import (
	"math"
	"time"

	"github.com/kvloginov/t3oe/internal/base"
)

type AIPlatformController struct {
	selfPosition           *base.Positional
	targetPosition         *base.Positional
	shootingAngleTolerance float64
	minDistance            float64
	maxDistance            float64

	// Decision making delay
	lastDecisionTime time.Time
	decisionDelay    time.Duration
	currentDecision  aiDecision
}

type aiDecision struct {
	forward bool
	left    bool
	right   bool
	shoot   bool
}

func NewAIPlatformController(selfPosition *base.Positional, targetPosition *base.Positional) *AIPlatformController {
	return &AIPlatformController{
		selfPosition:           selfPosition,
		targetPosition:         targetPosition,
		shootingAngleTolerance: math.Pi / 6,           // 30 degrees - less precise
		minDistance:            2.0,                   // much smaller minimum distance for shooting
		maxDistance:            25.0,                  // much larger maximum distance
		decisionDelay:          50 * time.Millisecond, // 50ms delay
		lastDecisionTime:       time.Now(),
		currentDecision:        aiDecision{},
	}
}

func (a *AIPlatformController) Forward() bool {
	a.updateDecisionIfNeeded()
	return a.currentDecision.forward
}

func (a *AIPlatformController) Backward() bool {
	// AI can't move backward - always return false
	return false
}

func (a *AIPlatformController) Left() bool {
	a.updateDecisionIfNeeded()
	return a.currentDecision.left
}

func (a *AIPlatformController) Right() bool {
	a.updateDecisionIfNeeded()
	return a.currentDecision.right
}

func (a *AIPlatformController) Shoot() bool {
	a.updateDecisionIfNeeded()
	return a.currentDecision.shoot
}

func (a *AIPlatformController) updateDecisionIfNeeded() {
	now := time.Now()
	if now.Sub(a.lastDecisionTime) < a.decisionDelay {
		return // Don't update decisions too frequently
	}

	a.lastDecisionTime = now
	a.currentDecision = a.makeDecision()
}

func (a *AIPlatformController) makeDecision() aiDecision {
	decision := aiDecision{}

	if a.selfPosition == nil || a.targetPosition == nil {
		return decision
	}

	distance := a.getDistanceToTarget()
	angleToTarget := a.getAngleToTarget()
	angleDiff := a.normalizeAngle(angleToTarget - a.selfPosition.Angle)

	// Movement decisions - keep larger distance for movement
	movementMinDistance := 8.0 // AI tries to stay at least 8 units away for movement

	if distance > a.maxDistance {
		decision.forward = true
	} else if distance > movementMinDistance && distance <= a.maxDistance {
		// Only move forward if roughly facing the target
		if math.Abs(angleDiff) < math.Pi/2 {
			decision.forward = true
		}
	}

	// Turning decisions with larger threshold to avoid jittering
	if angleDiff < -0.2 {
		decision.left = true
	} else if angleDiff > 0.2 {
		decision.right = true
	}

	// Shooting decisions - use smaller minimum distance for shooting
	if distance >= a.minDistance && distance <= a.maxDistance {
		if math.Abs(angleDiff) <= a.shootingAngleTolerance {
			decision.shoot = true
		}
	}
	return decision
}

// Helper functions

func (a *AIPlatformController) getDistanceToTarget() float64 {
	if a.selfPosition == nil || a.targetPosition == nil {
		return 0
	}

	diff := a.targetPosition.Pos.Minus(a.selfPosition.Pos)
	return diff.Length()
}

func (a *AIPlatformController) getAngleToTarget() float64 {
	if a.selfPosition == nil || a.targetPosition == nil {
		return 0
	}

	diff := a.targetPosition.Pos.Minus(a.selfPosition.Pos)
	return math.Atan2(diff.Y, diff.X)
}

func (a *AIPlatformController) normalizeAngle(angle float64) float64 {
	for angle > math.Pi {
		angle -= 2 * math.Pi
	}
	for angle < -math.Pi {
		angle += 2 * math.Pi
	}
	return angle
}

func (a *AIPlatformController) SetSelfPosition(pos *base.Positional) {
	a.selfPosition = pos
}

func (a *AIPlatformController) SetTargetPosition(pos *base.Positional) {
	a.targetPosition = pos
}
