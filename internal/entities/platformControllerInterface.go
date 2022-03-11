package entities

type PlatformController interface {
	Forward() bool
	Backward() bool
	Left() bool
	Right() bool
	Shoot() bool
}
