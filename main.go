package main

import (
	"github.com/sidav/sidavgorandom/fibrandom"
)

var io consoleIO
var rnd fibrandom.FibRandom

func main() {
	io.init()
	rnd.InitDefault()
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
		x = x/abs(x)
	}
	if y != 0 {
		y = y/abs(y)
	}
	return x, y
}
