package objects

import (
	"fmt"
	"littlejumbo/greak/values"

	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/physics"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

type Brick struct {
	instance *lifecycle.GameObject
	name     string
	rect     utils.RectSpecs
	body     physics.RigidBody
	color    render.Color
	point    int
	power    string
}

func NewBrick(name string, rect utils.RectSpecs, color render.Color) *Brick {
	brick := &Brick{
		name:  name,
		rect:  rect,
		color: color,
	}

	brick.instance = lifecycle.Register(&lifecycle.GameObject{
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

func (b *Brick) SetPowerUp(value string) {
	b.power = value
}

func (b *Brick) start() {
	b.body = physics.RegisterBody(&b.rect, b.name)

	if b.color == render.Green {
		events.Subscribe(values.COLOR_GREEN_DOUBLE_POINTS_EVENT, func(params ...any) error {
			b.doublePoints()
			return nil
		})
	}

	if b.color == render.Yellow {
		events.Subscribe(values.COLOR_YELLOW_DOUBLE_POINTS_EVENT, func(params ...any) error {
			b.doublePoints()
			return nil
		})
	}

	if b.color == render.Orange {
		events.Subscribe(values.COLOR_ORANGE_DOUBLE_POINTS_EVENT, func(params ...any) error {
			b.doublePoints()
			return nil
		})
	}

	if b.color == render.Red {
		events.Subscribe(values.COLOR_RED_DOUBLE_POINTS_EVENT, func(params ...any) error {
			b.doublePoints()
			return nil
		})
	}

	events.Subscribe(values.COLOR_ALL_DOUBLE_POINTS_EVENT, func(params ...any) error {
		b.doublePoints()
		return nil
	})
}

func (b *Brick) physics() {
	collision := physics.CheckCollision(&b.body)

	if collision.Name != "nil" {
		events.Emit(b.power)
		events.Emit(values.BRICK_DESTROYED_EVENT, b.name, b.rect, b.color, b.point)
		lifecycle.Stop(b.instance)
	}
}

func (b *Brick) render() {
	render.DrawSimpleShapes(b.rect, b.color)
}

func (b *Brick) destroy() {
	physics.RemoveBody(&b.body)
}

func (b *Brick) doublePoints() {
	b.point += b.point
	fmt.Printf("Doubling points for %v. Current: %v\n", b.color, b.point)
}
