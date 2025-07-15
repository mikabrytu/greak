package main

import (
	values "littlejumbo/greak/values"

	gomesengine "github.com/mikabrytu/gomes-engine"
	"github.com/mikabrytu/gomes-engine/lifecycle"
)

func main() {
	gomesengine.Init(
		values.APP_TITLE,
		int32(values.SCREEN_SIZE.X),
		int32(values.SCREEN_SIZE.Y),
	)
	lifecycle.SetSmoothStep(0.9)

	game()

	gomesengine.Run()
}
