package objects

import (
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

type Ball struct {
	rect  utils.RectSpecs
	color render.Color
}

func NewBall(rect utils.RectSpecs, color render.Color) *Ball {
	ball := &Ball{
		rect:  rect,
		color: color,
	}

	lifecycle.Register(lifecycle.GameObject{
		Render: ball.render,
	})

	return ball
}

func (b *Ball) render() {
	render.DrawSimpleShapes(b.rect, b.color)
}
