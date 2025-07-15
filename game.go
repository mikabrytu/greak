package main

import (
	"fmt"

	"littlejumbo/greak/managers/score"
	"littlejumbo/greak/objects"
	"littlejumbo/greak/values"

	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/ui"
	"github.com/mikabrytu/gomes-engine/utils"
)

var font *ui.Font

func game() {
	preparePaddle()
	prepareBall()
	prepareBricks()
	prepareScore()

	registerEvents()
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
	ball.SetDirection(math.Vector2{
		X: 0,
		Y: 1,
	})
}

func prepareBricks() {
	offset := 2
	startY := 50
	rect := utils.RectSpecs{
		Width:  values.BRICK_SIZE.X,
		Height: values.BRICK_SIZE.Y,
	}

	for i := 0; i < values.BRICK_ROWS; i++ {
		for j := 0; j < values.BRICK_COLS; j++ {
			rect.PosX = j * (values.BRICK_SIZE.X + offset)
			rect.PosY = startY + (i * (values.BRICK_SIZE.Y + offset))

			color := render.White
			point := 0
			switch i {
			case 0, 1:
				color = render.Red
				point = 7
			case 2, 3:
				color = render.Orange
				point = 5
			case 4, 5:
				color = render.Yellow
				point = 3
			case 6, 7:
				color = render.Green
				point = 1
			}

			name := fmt.Sprintf("brick-%v%v", i, j)
			brick := objects.NewBrick(name, rect, color)
			brick.SetPoint(point)
		}
	}
}

func prepareScore() {
	specs := ui.FontSpecs{
		Name: "Pixel",
		Path: "assets/font/CutePixel.ttf",
		Size: 42,
	}
	position := math.Vector2{
		X: 10,
		Y: 10,
	}

	score.Init()
	font = ui.NewFont(specs, values.SCREEN_SIZE)
	font.Init("0", render.White, position)
}

func registerEvents() {
	events.Subscribe(values.BRICK_DESTROYED_EVENT, func(params ...any) error {
		point := params[0].([]any)[0].([]any)[0]
		score.Add(point.(int))
		font.UpdateText(fmt.Sprintf("%d", score.Show()))

		return nil
	})
}
