package gameObjects

import (
	"errors"
	"github.com/google/uuid"
)

// GameObjects UVAGA global object
var GameObjects = newGameObjects()

var errGameObjectNameIsInvalid = errors.New("global object name is invalid")

type gameObjects struct {
	objects map[string]GameObject
}

func newGameObjects() *gameObjects {
	return &gameObjects{
		objects: make(map[string]GameObject, 0),
	}
}

// RegisterWithGeneratedId registers new global object with randomized name, returns its name in result and sets it by HasName interface
func (g *gameObjects) RegisterWithGeneratedId(object GameObject) string {
	uniqueKey := g.generateUniqueKey()
	object.SetName(uniqueKey)
	g.objects[uniqueKey] = object
	return uniqueKey
}

//// RegisterName registers new global object with specified name
//func (g *gameObjects) RegisterName(name string, object GameObject) error {
//	if len(name) <= 0 {
//		return fmt.Errorf("%w: name: %v", errGameObjectNameIsInvalid, name)
//	}
//	g.objects[name] = object
//	return nil
//}

func (g *gameObjects) GetAll() map[string]GameObject {
	return g.objects
}

func (g *gameObjects) generateUniqueKey() string {
	// we assume that the uuid is always unique
	return uuid.New().String()
}

func (g *gameObjects) Destroy(name string) {
	delete(g.objects, name)
}
