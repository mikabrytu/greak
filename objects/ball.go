package objects

import (
	"fmt"
	"littlejumbo/greak/values"

	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/physics"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

type Ball struct {
	name      string
	rect      utils.RectSpecs
	body      physics.RigidBody
	color     render.Color
	speed     int
	direction math.Vector2
	inPlay    bool
}

func NewBall(name string, rect utils.RectSpecs, color render.Color) *Ball {
	ball := &Ball{
		name:   name,
		rect:   rect,
		color:  color,
		inPlay: true,
	}

	lifecycle.Register(&lifecycle.GameObject{
		Start:   ball.start,
		Physics: ball.physics,
		Render:  ball.render,
	})

	return ball
}

func (b *Ball) SetSpeed(speed int) {
	b.speed = speed
}

func (b *Ball) SetDirection(direction math.Vector2) {
	b.direction = direction
	if b.body.Name != "" && b.body.Name != "nil" {
		b.body.Axis = b.direction
	}
}

func (b *Ball) SetPosition(position math.Vector2) {
	b.rect.PosX = position.X
	b.rect.PosY = position.Y
}

func (b *Ball) SetInPlay(inPlay bool) {
	b.inPlay = inPlay
}

func (b *Ball) start() {
	b.body = physics.RegisterBody(&b.rect, b.name)
	b.body.Axis = b.direction

	physics.EnableDynamicCollision(&b.body)

	events.Subscribe(values.BALL_INCREASE_SPEED_EVENT, func(params ...any) error {
		b.speed += values.BALL_SPEED_INCREMENT

		fmt.Printf("Ball Speed increased to %v\n", b.speed)

		return nil
	})
}

func (b *Ball) physics() {
	b.rect.PosX += b.body.Axis.X * b.speed
	b.rect.PosY += b.body.Axis.Y * b.speed

	physics.ResolveDynamicCollisions(&b.body, false, false)

	b.checkScreenBoundaries()
}

func (b *Ball) render() {
	render.DrawSimpleShapes(b.rect, b.color)
}

func (b *Ball) checkScreenBoundaries() {
	if b.body.Rect.PosX < 0 {
		b.body.Axis.X = 1
	}

	if b.body.Rect.PosY < 0 {
		b.body.Axis.Y = 1
	}

	if (b.body.Rect.PosX + b.body.Rect.Width) > values.SCREEN_SIZE.X {
		b.body.Axis.X = -1
	}

	if b.inPlay && (b.body.Rect.PosY+b.body.Rect.Height) > values.SCREEN_SIZE.Y {
		b.inPlay = false
		events.Emit(values.BALL_OUT_EVENT)
	}
}
