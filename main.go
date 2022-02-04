package main

import (
	"github.com/sidav/sidavgorandom/fibrandom"
)

var r cliIO
var rnd fibrandom.FibRandom

func main() {
	r.init()
	rnd.InitDefault()
	runGame()
	r.close()
}
