package game

import (
	"container/list"
	"fmt"
	"littlejumbo/greak/managers/score"
	"littlejumbo/greak/objects"
	"littlejumbo/greak/values"
	"math/rand"

	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/ui"
	"github.com/mikabrytu/gomes-engine/utils"
)

var powerList *list.List

func preparePaddle() {
	rect := utils.RectSpecs{
		PosX:   (values.SCREEN_SIZE.X / 2) - (values.BRICK_SIZE.X / 2),
		PosY:   values.SCREEN_SIZE.Y - PADDLE_OFFSET,
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

	// Generate 22 random ids for power ups
	powerList = list.New()
	for i := 0; i < 22; i++ {
		id := math.Vector2{
			X: rand.Intn(values.BRICK_ROWS),
			Y: rand.Intn(values.BRICK_COLS),
		}

		powerList.PushBack(id)
	}

	for i := range values.BRICK_ROWS {
		for j := range values.BRICK_COLS {
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

			for e := powerList.Front(); e != nil; e = e.Next() {
				id := e.Value.(math.Vector2)
				if id.X == i && id.Y == j {
					brick.SetPowerUp(100)
					break
				}
			}
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
	scoreText = ui.NewFont(specs, values.SCREEN_SIZE)
	scoreText.Init("0", render.White, position)

	position.X = values.SCREEN_SIZE.X / 2
	text := fmt.Sprintf("%d", score.ShowHigh())
	highText = ui.NewFont(specs, values.SCREEN_SIZE)
	highText.Init(text, render.White, position)
}
