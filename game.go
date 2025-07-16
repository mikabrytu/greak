package main

import (
	"container/list"
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

type BrickData struct {
	Name  string
	Rect  utils.RectSpecs
	Color render.Color
	Value int
}

var paddle *objects.Paddle
var ball *objects.Ball
var scoreText *ui.Font
var highText *ui.Font
var destroyedBricks *list.List

const PADDLE_OFFSET int = 24

func game() {
	destroyedBricks = list.New()
	score.Init()

	preparePaddle()
	prepareBall()
	prepareBricks()
	prepareScore()

	registerEvents()
}

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
