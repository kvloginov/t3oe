# T3OE Game Architecture Documentation

## Overview

T3OE is a 2D duel game built using the Ebiten game engine in Go. The architecture follows a component-based design pattern with a global game object system, physics simulation, and trigger-based collision detection.

## Core Architecture

### 1. Game Loop Structure

The main game loop is implemented in `internal/game.go`:

```go
func (g *Game) Update() error {
    dt := float64(1) / 60
    objects := gameObjects.GameObjects.GetAll()
    for i, _ := range objects {
        objects[i].Update(dt)
    }
    trigger.CheckTriggers()
    return nil
}
```

- **Update Phase**: All game objects are updated with delta time
- **Physics Simulation**: Objects with physics components are moved and rotated
- **Collision Detection**: Trigger system checks for collisions between objects
- **Rendering**: All drawable objects are rendered to the screen

### 2. Entity Component System

The game uses a composition-based approach where entities are composed of multiple components:

#### Base Components

- **`base.Positional`**: Position and rotation in world space
- **`base.Physical`**: Physics properties (velocity, acceleration, drag)
- **`base.VolumeObject`**: Collision bounds and pivot point
- **`entities.NamedEntity`**: Unique identifier for the entity
- **`entities.Team`**: Team affiliation (BLUE, RED, UNDEFINED)

#### Core Interfaces

- **`gameObjects.GameObject`**: Main interface combining Drawable, Updatable, and HasName
- **`gameObjects.Drawable`**: Entities that can be rendered
- **`gameObjects.Updatable`**: Entities that update each frame
- **`gameObjects.HasName`**: Entities with unique identifiers

### 3. Global Game Object System

The `gameObjects.GameObjects` singleton manages all active game entities:

```go
var GameObjects = newGameObjects()

func (g *gameObjects) RegisterWithGeneratedId(object GameObject) string
func (g *gameObjects) GetAll() map[string]GameObject
func (g *gameObjects) Destroy(name string)
```

**Key Features:**
- Automatic UUID generation for entity names
- Centralized entity lifecycle management
- Sorted rendering order based on entity names

### 4. Physics System

Physics simulation is handled by the `base.Physical` component:

- **Position Integration**: Uses Verlet integration for position updates
- **Velocity and Acceleration**: Applied with delta time
- **Drag**: Air resistance based on velocity squared
- **Rotation**: Angular velocity with resistance coefficient

### 5. Trigger System

Collision detection uses a trigger-based approach in `internal/trigger/`:

```go
func RegisterTrigger(owner TriggerOwner, hitboxRadius float64, onHitAction func(object1 interface{}, object2 interface{}))
func CheckTriggers() // O(n²) collision detection
```

**Features:**
- Circular hitboxes with configurable radius
- Callback-based collision response
- Automatic cleanup when entities are destroyed

## Entity Creation Guide

### Step 1: Define Entity Structure

Create a new entity by composing base components:

```go
type MyEntity struct {
    base.Physical        // Physics simulation
    base.VolumeObject   // Collision bounds
    entities.Team       // Team affiliation
    entities.NamedEntity // Unique identifier
    
    // Custom fields
    customField string
    image       *ebiten.Image
}
```

### Step 2: Constructor Function

Create a constructor that initializes the entity and registers it:

```go
func NewMyEntity(position base.Positional, team entities.Team) *MyEntity {
    entity := &MyEntity{
        Physical: base.Physical{
            Positional: position,
            DragCoefficient: 0.1,
        },
        VolumeObject: base.VolumeObject{
            PivotRelativeX: 0.5,
            PivotRelativeY: 0.5,
            Width:          1.0,
            Height:         1.0,
        },
        Team: team,
        // Initialize custom fields
        customField: "example",
        image:       drawing.SOME_IMAGE,
    }
    
    // Register with global system
    entity.name = gameObjects.GameObjects.RegisterWithGeneratedId(entity)
    
    // Register trigger if needed
    trigger.RegisterTrigger(entity, 1.0, entity.onCollision)
    
    return entity
}
```

### Step 3: Implement Required Interfaces

#### Update Method (Updatable)

```go
func (e *MyEntity) Update(dt float64) {
    // Custom update logic
    e.customField = "updated"
    
    // Apply physics if needed
    e.Physical.Update(dt)
}
```

#### Draw Method (Drawable)

```go
func (e *MyEntity) Draw(screen *ebiten.Image, drawingStuff *drawing.DrawingStuff) {
    // Render the entity
    drawingStuff.DrawVolumeObject(e.VolumeObject, e.Positional, e.image, screen)
    
    // Optional: Draw debug information
    drawingStuff.DrawDebugPositionPoint(e.Positional, screen)
}
```

### Step 4: Collision Handling

If your entity needs collision detection:

```go
func (e *MyEntity) onCollision(object1 interface{}, object2 interface{}) {
    switch other := object2.(type) {
    case entities.HasTeam:
        if other.GetTeam() != e.GetTeam() {
            // Handle collision with enemy
            e.handleEnemyCollision(other)
        }
    case *entities.Bullet:
        // Handle collision with bullet
        e.handleBulletCollision(other)
    }
}
```

## File Organization

### Directory Structure

```
internal/
├── base/                    # Base components
│   ├── positional.go       # Position and rotation
│   ├── physical.go         # Physics simulation
│   ├── vector.go           # Vector math utilities
│   └── volumeObject.go     # Collision bounds
├── entities/               # Game entities
│   ├── platform.go         # Player platforms
│   ├── bullet.go           # Projectiles
│   ├── explosion.go        # Visual effects
│   ├── gun.go              # Weapon systems
│   ├── team.go             # Team definitions
│   └── namedEntity.go      # Name component
├── gameObjects/            # Entity management
│   ├── gameObjects.go      # Global entity system
│   └── *Interface.go       # Core interfaces
├── controllers/            # Input handling
│   ├── directInputPlatformController.go
│   └── randomController.go
├── drawing/                # Rendering system
│   └── drawingStuff.go     # Drawing utilities
└── trigger/                # Collision system
    ├── trigger.go          # Trigger management
    └── hasPositionInterface.go
```

### Where to Place New Entities

1. **Simple Entities**: Place in `internal/entities/` directory
2. **Complex Systems**: Create a new subdirectory under `internal/`
3. **Interfaces**: Define in the relevant package (e.g., `gameObjects/` for core interfaces)
4. **Controllers**: Place in `internal/controllers/` for input handling

## Entity Interaction Patterns

### 1. Direct Component Access

Entities can directly access each other's components:

```go
if platform, ok := other.(*entities.Platform); ok {
    position := platform.GetPosition()
    team := platform.GetTeam()
}
```

### 2. Interface-Based Interaction

Use interfaces for loose coupling:

```go
if hasTeam, ok := other.(entities.HasTeam); ok {
    if hasTeam.GetTeam() == e.GetTeam() {
        // Same team logic
    }
}
```

### 3. Trigger-Based Events

Use the trigger system for collision-based interactions:

```go
trigger.RegisterTrigger(entity, radius, func(obj1, obj2 interface{}) {
    // Collision response
})
```

### 4. Global Systems

Access global systems through singleton patterns:

```go
// Spawn new entity
newEntity := entities.NewBullet(position, team)

// Destroy entity
gameObjects.GameObjects.Destroy(entityName)
```

## Best Practices

### 1. Entity Lifecycle

- Always register entities with `gameObjects.GameObjects.RegisterWithGeneratedId()`
- Clean up triggers when destroying entities: `trigger.DeleteTriggers(entityName)`
- Use proper constructor functions for initialization

### 2. Performance Considerations

- The trigger system is O(n²) - avoid too many collision objects
- Use appropriate drag coefficients to prevent physics instability
- Consider object pooling for frequently created/destroyed entities

### 3. Team System

- Use the `entities.Team` type for team affiliation
- Implement `entities.HasTeam` interface for team-based interactions
- Check team affiliation in collision handlers

### 4. Physics Integration

- Always call `Physical.Update(dt)` in your entity's Update method
- Set appropriate drag coefficients to prevent runaway physics
- Use turning resistance for smooth rotation

### 5. Rendering

- Implement proper pivot points in `VolumeObject`
- Use the `drawing.DrawingStuff` utilities for consistent rendering
- Consider z-ordering through entity naming conventions

## Example: Adding a New Projectile

```go
// internal/entities/rocket.go
package entities

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/kvloginov/t3oe/internal/base"
    "github.com/kvloginov/t3oe/internal/drawing"
    "github.com/kvloginov/t3oe/internal/gameObjects"
    "github.com/kvloginov/t3oe/internal/trigger"
)

type Rocket struct {
    base.Physical
    base.VolumeObject
    entities.Team
    entities.NamedEntity
    
    fuel     float64
    thrust   float64
}

func NewRocket(pos base.Positional, team Team, thrust float64) *Rocket {
    rocket := &Rocket{
        Physical: base.Physical{
            Positional: pos,
            DragCoefficient: 0.05,
        },
        VolumeObject: base.VolumeObject{
            PivotRelativeX: 0.5,
            PivotRelativeY: 0.5,
            Width:          0.5,
            Height:         1.0,
        },
        Team:   team,
        fuel:   10.0,
        thrust: thrust,
    }
    
    id := gameObjects.GameObjects.RegisterWithGeneratedId(rocket)
    trigger.RegisterTrigger(rocket, 0.5, rocket.onCollision)
    
    return rocket
}

func (r *Rocket) Update(dt float64) {
    if r.fuel > 0 {
        // Apply thrust
        thrustVector := base.NewVectorWithAngle(r.Angle).MultiplyScalar(r.thrust)
        r.Acceleration = r.Acceleration.Plus(thrustVector)
        r.fuel -= dt
    }
    
    r.Physical.Update(dt)
}

func (r *Rocket) Draw(screen *ebiten.Image, drawingStuff *drawing.DrawingStuff) {
    image := drawing.ROCKET_BLUE_IMG
    if r.Team == TEAM_RED {
        image = drawing.ROCKET_RED_IMG
    }
    
    drawingStuff.DrawVolumeObject(r.VolumeObject, r.Positional, image, screen)
}

func (r *Rocket) onCollision(obj1, obj2 interface{}) {
    if hasTeam, ok := obj2.(HasTeam); ok {
        if hasTeam.GetTeam() != r.GetTeam() {
            // Create explosion
            NewExplosion(r.Positional)
            
            // Destroy rocket
            gameObjects.GameObjects.Destroy(r.GetName())
            trigger.DeleteTriggers(r.GetName())
        }
    }
}
```

This architecture provides a flexible foundation for adding new game entities while maintaining clean separation of concerns and efficient performance. 