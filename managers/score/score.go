package score

var score int

func Init() {
	score = 0
}

func Add(point int) {
	score += point
}

func Show() int {
	return score
}

func Reset() {
	score = 0
}
