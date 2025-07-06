package objects

import (
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
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

func NewPaddle(rect utils.RectSpecs, color render.Color) *Paddle {
	paddle := &Paddle{
		rect:    rect,
		color:   color,
		canMove: true,
	}

	lifecycle.Register(lifecycle.GameObject{
		Start:   paddle.start,
		Physics: paddle.physics,
		Render:  paddle.render,
	})

	return paddle
}

func (p *Paddle) SetSpeed(speed int) {
	p.speed = speed
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
}

func (p *Paddle) physics() {
	// TODO: Block movement when reach screen boundaries

	if p.canMove {
		p.rect.PosX += p.axis * p.speed
	}
}

func (p *Paddle) render() {
	render.DrawSimpleShapes(p.rect, p.color)
}

func (p *Paddle) move(axis int) {
	p.axis = axis
}
