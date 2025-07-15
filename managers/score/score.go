package score

import (
	"os"
	"time"

	savesystem "github.com/mikabrytu/gomes-engine/systems/save"
)

type Score struct {
	Score     int       `json:"score"`
	Timestamp time.Time `json:"timestamp"`
}

var current int
var high int
var file string
var score Score

func Init() {
	current = 0
	file = "highscore.json"

	if !fileExists(file) {
		createFile(file)
	}

	err := savesystem.Load(file, &score)
	if err != nil {
		panic(err)
	}

	high = score.Score
}

func Add(point int) {
	current += point
}

func ShowCurrent() int {
	return current
}

func ShowHigh() int {
	return high
}

func SaveCurrent() {
	if current <= high {
		return
	}

	high = current
	newScore := Score{
		Score:     high,
		Timestamp: time.Now(),
	}
	savesystem.Save(newScore, file)
}

func Reset() {
	current = 0
}

func createFile(path string) {
	var empty Score
	savesystem.Save(empty, path)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}
