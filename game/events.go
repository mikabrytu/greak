package game

import (
	"fmt"
	"littlejumbo/greak/managers/score"
	"littlejumbo/greak/values"

	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

func registerEvents() {
	events.Subscribe(values.BRICK_DESTROYED_EVENT, func(params ...any) error {
		raw := params[0].([]any)[0].([]any)
		data := BrickData{
			Name:  raw[0].(string),
			Rect:  raw[1].(utils.RectSpecs),
			Color: raw[2].(render.Color),
			Value: raw[3].(int),
		}

		onBrickDestroyed(data)
		return nil
	})

	events.Subscribe(values.BALL_OUT_EVENT, func(params ...any) error {
		onBallOut()
		return nil
	})
}

func onBrickDestroyed(data BrickData) {
	destroyedBricks.PushBack(data)

	score.Add(data.Value)
	scoreText.UpdateText(fmt.Sprintf("%d", score.ShowCurrent()))
}

func onBallOut() {
	updateScore()
	resetBall()
	resetPaddle()
	resetBricks()
}
