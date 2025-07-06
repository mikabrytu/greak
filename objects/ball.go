package objects

import (
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/physics"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

type Ball struct {
	name  string
	rect  utils.RectSpecs
	color render.Color
	speed int
	body  physics.RigidBody
}

func NewBall(name string, rect utils.RectSpecs, color render.Color) *Ball {
	ball := &Ball{
		name:  name,
		rect:  rect,
		color: color,
	}

	lifecycle.Register(lifecycle.GameObject{
		Start:   ball.start,
		Physics: ball.physics,
		Render:  ball.render,
	})

	return ball
}

func (b *Ball) SetSpeed(speed int) {
	b.speed = speed
}

func (b *Ball) start() {
	b.body = physics.RegisterBody(&b.rect, b.name)
	b.body.Axis.Y = -1
	physics.EnableDynamicCollision(&b.body)
}

func (b *Ball) physics() {
	b.rect.PosX += b.body.Axis.X * b.speed
	b.rect.PosY += b.body.Axis.Y * b.speed

	physics.ResolveDynamicCollisions(&b.body)
}

func (b *Ball) render() {
	render.DrawSimpleShapes(b.rect, b.color)
}
