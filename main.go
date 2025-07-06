package main

import (
	"fmt"
	"littlejumbo/greak/objects"
	values "littlejumbo/greak/values"

	gomesengine "github.com/mikabrytu/gomes-engine"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

var paddle *objects.Paddle
var ball *objects.Ball

func main() {
	gomesengine.Init(
		values.APP_TITLE,
		int32(values.SCREEN_SIZE.X),
		int32(values.SCREEN_SIZE.Y),
	)

	game()

	gomesengine.Run()
}

func game() {
	preparePaddle()
	prepareBall()
	prepareBricks()
	prepareScore()
}

func preparePaddle() {
	offset := 24
	rect := utils.RectSpecs{
		PosX:   (values.SCREEN_SIZE.X / 2) - (values.BRICK_SIZE.X / 2),
		PosY:   values.SCREEN_SIZE.Y - offset,
		Width:  values.BRICK_SIZE.X,
		Height: values.BRICK_SIZE.Y,
	}

	paddle = objects.NewPaddle("paddle", rect, render.Blue)
	paddle.SetSpeed(values.PADDLE_SPEED)
}

func prepareBall() {
	rect := utils.RectSpecs{
		PosX:   (values.SCREEN_SIZE.X / 2) - (values.BALL_SIZE.X / 2),
		PosY:   (values.SCREEN_SIZE.Y / 2) - (values.BALL_SIZE.Y / 2),
		Width:  values.BALL_SIZE.X,
		Height: values.BALL_SIZE.Y,
	}

	ball = objects.NewBall("ball", rect, render.White)
	ball.SetSpeed(values.BALL_SPEED)
}

func prepareBricks() {
	offset := 2
	rect := utils.RectSpecs{
		Width:  values.BRICK_SIZE.X,
		Height: values.BRICK_SIZE.Y,
	}

	for i := 0; i < values.BRICK_ROWS; i++ {
		for j := 0; j < values.BRICK_COLS; j++ {
			rect.PosX = j * (values.BRICK_SIZE.X + offset)
			rect.PosY = i * (values.BRICK_SIZE.Y + offset)

			color := render.White
			switch i {
			case 0, 1:
				color = render.Red
			case 2, 3:
				color = render.Orange
			case 4, 5:
				color = render.Yellow
			case 6, 7:
				color = render.Green
			}

			name := fmt.Sprintf("brick-%v%v", i, j)
			objects.NewBrick(name, rect, color)
		}
	}
}

func prepareScore() {

}
