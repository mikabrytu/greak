package objects

import (
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/physics"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

type Brick struct {
	name  string
	rect  utils.RectSpecs
	color render.Color
}

func NewBrick(name string, rect utils.RectSpecs, color render.Color) *Brick {
	brick := &Brick{
		name:  name,
		rect:  rect,
		color: color,
	}

	lifecycle.Register(lifecycle.GameObject{
		Start:  brick.start,
		Render: brick.render,
	})

	return brick
}

func (b *Brick) start() {
	physics.RegisterBody(&b.rect, b.name)
}

func (b *Brick) render() {
	render.DrawSimpleShapes(b.rect, b.color)
}
