package objects

import (
	"littlejumbo/greak/values"

	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/physics"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

type Brick struct {
	instance lifecycle.GameObject
	name     string
	rect     utils.RectSpecs
	body     physics.RigidBody
	color    render.Color
	point    int
}

func NewBrick(name string, rect utils.RectSpecs, color render.Color) *Brick {
	brick := &Brick{
		name:  name,
		rect:  rect,
		color: color,
	}

	brick.instance = lifecycle.Register(lifecycle.GameObject{
		Start:   brick.start,
		Physics: brick.physics,
		Render:  brick.render,
		Destroy: brick.destroy,
	})

	return brick
}

func (b *Brick) SetPoint(value int) {
	b.point = value
}

func (b *Brick) start() {
	b.body = physics.RegisterBody(&b.rect, b.name)
}

func (b *Brick) physics() {
	collision := physics.CheckCollision(&b.body)

	if collision.Name != "nil" {
		events.Emit(values.BRICK_DESTROYED_EVENT, b.point)
		lifecycle.Stop(b.instance)
	}
}

func (b *Brick) render() {
	render.DrawSimpleShapes(b.rect, b.color)
}

func (b *Brick) destroy() {
	physics.RemoveBody(&b.body)
}
