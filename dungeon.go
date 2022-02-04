package main

import (
	"github.com/sidav/cyclicdungeongenerator/generator"
)

type dungeon struct {
	name           string
	layout         generator.LayoutInterface
	rooms          [][]*room
	totalStages    int
}

func (d *dungeon) initAndGenerate(patternFileName string) (int, int) {
	const sizew, sizeh = 6, 4
	playerx, playery := 0, 0
	gen := generator.InitGeneratorsWrapper()
	d.layout, _ = gen.GenerateLayout(sizew, sizeh, patternFileName)
	d.rooms = make([][]*room, sizew)
	for i := range d.rooms {
		d.rooms[i] = make([]*room, sizeh)
	}
	for x := range d.rooms {
		for y := range d.rooms[x] {
			if d.layout.GetElement(x, y).IsNode() {
				d.rooms[x][y] = &room{
					isCleared:   false,
					isGenerated: false,
				}
			}
			if d.layout.GetElement(x, y).IsNode() && d.layout.GetElement(x, y).HasTag("start") {
				playerx = x
				playery = y
			}
		}
	}
	return playerx, playery
}
