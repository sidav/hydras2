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
