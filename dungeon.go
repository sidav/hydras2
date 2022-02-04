package main

import (
	"github.com/sidav/cyclicdungeongenerator/generator"
)

type dungeon struct {
	name   string
	layout generator.LayoutInterface
	rooms  [][]*dungeonCell
}

func (d *dungeon) initAndGenerate(patternFileName string) (int, int) {
	const sizew, sizeh = 7, 4
	playerx, playery := 0, 0
	gen := generator.InitGeneratorsWrapper()
	d.layout, _ = gen.GenerateLayout(sizew, sizeh, patternFileName)
	d.rooms = make([][]*dungeonCell, sizew)
	for i := range d.rooms {
		d.rooms[i] = make([]*dungeonCell, sizeh)
	}
	for x := range d.rooms {
		for y := range d.rooms[x] {
			d.rooms[x][y] = &dungeonCell{
				isCleared:   false,
				isRoom:      d.layout.GetElement(x, y).IsNode(),
			}
			if d.layout.GetElement(x, y).IsNode() && d.layout.GetElement(x, y).HasTag("start") {
				d.rooms[x][y].isCleared = true
				playerx = x
				playery = y
			}
		}
	}
	return playerx, playery
}

func (d *dungeon) canPlayerMoveFromByVector(p *player, vx, vy int) bool {
	conns := d.layout.GetElement(p.dungX, p.dungY).GetAllConnectionsCoords()
	for c := 0; c < len(conns); c++ {
		if conns[c][0] == vx && conns[c][1] == vy {
			return true
		}
	}
	return false
}
