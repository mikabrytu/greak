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
)

const (
	BRICK_ROWS int = 8
	BRICK_COLS int = 14
)
