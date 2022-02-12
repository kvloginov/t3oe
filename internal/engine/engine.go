package engine

type engine struct {
	gameObjects map[string]interface{} // object name to game object
}

func (e *engine) RegisterObject(name string, object interface{}) {
	e.gameObjects[name] = object
}

func (e *engine) GetObject(name string) interface{} {
	return e.gameObjects[name]
}

func (e *engine) Update() {
	//for key, _ := range e.gameObjects {
	//	obj := e.gameObjects[key]
	//	switch obj.(type) {
	//	case base.Physical:
	//		obj.SpeedX = 0
	//	case base.VolumeObject:
	//	case base.Positional:
	//	}
	//
	//}
}
