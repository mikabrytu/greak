package objects

import (
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

type Brick struct {
	rect  utils.RectSpecs
	color render.Color
}

func NewBrick(rect utils.RectSpecs, color render.Color) *Brick {
	brick := &Brick{
		rect:  rect,
		color: color,
	}

	lifecycle.Register(lifecycle.GameObject{
		Render: brick.render,
	})

	return brick
}

func (b *Brick) render() {
	render.DrawSimpleShapes(b.rect, b.color)
}
