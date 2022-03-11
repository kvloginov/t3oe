package trigger

import "github.com/kvloginov/t3oe/internal/base"

type TriggerOwner interface {
	GetPosition() base.Vector
	GetName() string
}
