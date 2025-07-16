package objects

import (
	"littlejumbo/greak/values"

	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/physics"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

type Paddle struct {
	name    string
	rect    utils.RectSpecs
	color   render.Color
	axis    int
	speed   int
	canMove bool
}

func NewPaddle(name string, rect utils.RectSpecs, color render.Color) *Paddle {
	paddle := &Paddle{
		name:    name,
		rect:    rect,
		color:   color,
		canMove: true,
	}

	lifecycle.Register(&lifecycle.GameObject{
		Start:   paddle.start,
		Physics: paddle.physics,
		Render:  paddle.render,
	})

	return paddle
}

func (p *Paddle) SetSpeed(speed int) {
	p.speed = speed
}

func (p *Paddle) SetPosition(position math.Vector2) {
	p.rect.PosX = position.X
	p.rect.PosY = position.Y
}

func (p *Paddle) SetSize(size math.Vector2) {
	p.rect.Width = size.X
	p.rect.Height = size.Y
}

func (p *Paddle) start() {
	physics.RegisterBody(&p.rect, p.name)

	events.Subscribe(events.INPUT_KEYBOARD_PRESSED_A, func(params ...any) error {
		p.move(-1)
		return nil
	})

	events.Subscribe(events.INPUT_KEYBOARD_RELEASED_A, func(params ...any) error {
		p.move(0)
		return nil
	})

	events.Subscribe(events.INPUT_KEYBOARD_PRESSED_D, func(params ...any) error {
		p.move(1)
		return nil
	})

	events.Subscribe(events.INPUT_KEYBOARD_RELEASED_D, func(params ...any) error {
		p.move(0)
		return nil
	})

	events.Subscribe(values.PADDLE_INCREASE_SIZE_EVENT, func(params ...any) error {
		p.rect.Width += values.PADDLE_SIZE_INCREMENT
		return nil
	})

	events.Subscribe(values.PADDLE_INCREASE_SPEED_EVENT, func(params ...any) error {
		p.speed += values.PADDLE_SPEED_INCREMENT
		return nil
	})
}

func (p *Paddle) physics() {
	if p.canMove {
		p.rect.PosX += p.axis * p.speed
	}
}

func (p *Paddle) render() {
	render.DrawSimpleShapes(p.rect, p.color)
}

func (p *Paddle) move(axis int) {
	if axis == -1 && p.rect.PosX < 0 {
		axis = 0
	}

	if axis == 1 && (p.rect.PosX+p.rect.Width) > values.SCREEN_SIZE.X {
		axis = 0
	}

	p.axis = axis
}
