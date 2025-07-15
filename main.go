package main

import (
	values "littlejumbo/greak/values"

	gomesengine "github.com/mikabrytu/gomes-engine"
)

func main() {
	gomesengine.Init(
		values.APP_TITLE,
		int32(values.SCREEN_SIZE.X),
		int32(values.SCREEN_SIZE.Y),
	)

	game()

	gomesengine.Run()
}
