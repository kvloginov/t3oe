package controllers

type emptyPlatformController struct {
}

func NewEmptyPlatformController() *emptyPlatformController {
	return &emptyPlatformController{}
}

func (c *emptyPlatformController) Forward() bool {
	return false
}

func (c *emptyPlatformController) Backward() bool {
	return false
}

func (c *emptyPlatformController) Left() bool {
	return false
}

func (c *emptyPlatformController) Right() bool {
	return false
}

func (c *emptyPlatformController) Shoot() bool {
	return false
}
