package main

import (
	"github.com/sidav/sidavgorandom/fibrandom"
	"hydras2/game_log"
)

var io consoleIO
var rnd *fibrandom.FibRandom
var log *game_log.GameLog
var exitGame bool

func main() {
	io.init()

	rnd = &fibrandom.FibRandom{}
	rnd.InitDefault()

	log = &game_log.GameLog{}
	log.Init(3)

	runGame()
	io.close()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func vectorToUnitVector(x, y int) (int, int) {
	if x != 0 {
		x = x / abs(x)
	}
	if y != 0 {
		y = y / abs(y)
	}
	return x, y
}
