package values

import "github.com/mikabrytu/gomes-engine/math"

var (
	SCREEN_SIZE = math.Vector2{
		X: 474,
		Y: 640,
	}
	BRICK_SIZE = math.Vector2{
		X: 32,
		Y: 16,
	}
	BALL_SIZE = math.Vector2{
		X: 12,
		Y: 12,
	}
	PADDLE_SIZE_INCREMENT int = BRICK_SIZE.X / 4
)

const (
	BRICK_ROWS             int = 8
	BRICK_COLS             int = 14
	PADDLE_SPEED           int = 5
	PADDLE_SPEED_INCREMENT int = 3
	BALL_SPEED             int = 3
)
