package game

import (
	"container/list"

	"littlejumbo/greak/managers/score"
	"littlejumbo/greak/objects"

	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/ui"
	"github.com/mikabrytu/gomes-engine/utils"
)

type BrickData struct {
	Name  string
	Rect  utils.RectSpecs
	Color render.Color
	Value int
}

type PowerUp struct {
	Id    math.Vector2
	Event string
}

var paddle *objects.Paddle
var ball *objects.Ball
var scoreText *ui.Font
var highText *ui.Font
var destroyedBricks *list.List
var powerList *list.List

const PADDLE_OFFSET int = 24

func Main() {
	destroyedBricks = list.New()
	score.Init()

	/* Setup */
	preparePaddle()
	prepareBall()
	prepareBricks()
	prepareScore()
	/* Setup */

	/* Events */
	registerEvents()
	/* Events */
}
