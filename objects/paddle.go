package objects

import (
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

type Paddle struct {
	rect  utils.RectSpecs
	color render.Color
}

func NewPaddle(rect utils.RectSpecs, color render.Color) *Paddle {
	paddle := &Paddle{
		rect:  rect,
		color: color,
	}

	lifecycle.Register(lifecycle.GameObject{
		Render: paddle.render,
	})

	return paddle
}

func (p *Paddle) render() {
	render.DrawSimpleShapes(p.rect, p.color)
}
