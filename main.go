package main

import (
	"littlejumbo/greak/objects"
	values "littlejumbo/greak/values"

	gomesengine "github.com/mikabrytu/gomes-engine"
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
