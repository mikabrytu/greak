package game

import (
	"container/list"
	"fmt"
	"littlejumbo/greak/managers/score"
	"littlejumbo/greak/objects"
	"littlejumbo/greak/values"

	"github.com/mikabrytu/gomes-engine/math"
)

func updateScore() {
	score.SaveCurrent()
	score.Reset()
	scoreText.UpdateText(fmt.Sprintf("%d", score.ShowCurrent()))
	highText.UpdateText(fmt.Sprintf("%d", score.ShowHigh()))
}

func resetBall() {
	bPos := math.Vector2{
		X: (values.SCREEN_SIZE.X / 2) - (values.BALL_SIZE.X / 2),
		Y: (values.SCREEN_SIZE.Y / 2) - (values.BALL_SIZE.Y / 2),
	}
	bDir := math.Vector2{
		X: 0,
		Y: 1,
	}

	ball.SetInPlay(true)
	ball.SetPosition(bPos)
	ball.SetSpeed(values.BALL_SPEED)
	ball.SetDirection(bDir)
}

func resetPaddle() {
	pPos := math.Vector2{
		X: (values.SCREEN_SIZE.X / 2) - (values.BRICK_SIZE.X / 2),
		Y: values.SCREEN_SIZE.Y - PADDLE_OFFSET,
	}

	paddle.SetPosition(pPos)
	paddle.SetSpeed(values.PADDLE_SPEED)
	paddle.SetSize(values.BRICK_SIZE)
}

func resetBricks() {
	if destroyedBricks.Len() == 0 {
		return
	}

	fmt.Printf("Bricks destroyed: %v\n", destroyedBricks.Len())

	for e := destroyedBricks.Front(); e != nil; e = e.Next() {
		data := e.Value.(BrickData)
		brick := objects.NewBrick(data.Name, data.Rect, data.Color)
		brick.SetPoint(data.Value)
	}

	destroyedBricks = list.New()
}
