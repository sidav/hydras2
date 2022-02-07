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
				isRoom: d.layout.GetElement(x, y).IsNode(),
			}
			if d.layout.GetElement(x, y).IsNode() && d.layout.GetElement(x, y).HasTag("start") {
				d.rooms[x][y].isGenerated = true
				playerx = x
				playery = y
			}
			if d.layout.GetElement(x, y).IsNode() && d.layout.GetElement(x, y).HasTag("ky1") {
				d.rooms[x][y].hasKey = 1
			}
			if d.layout.GetElement(x, y).IsNode() && d.layout.GetElement(x, y).HasTag("ky2") {
				d.rooms[x][y].hasKey = 2
			}
		}
	}
	return playerx, playery
}

func (d *dungeon) generateAndRevealRoomsAroundPlayer(p *player) {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if (x == 0 || y == 0) && d.canPlayerMoveFromByVector(p, x, y) {
				rx, ry := p.dungX+x, p.dungY+y
				d.rooms[rx][ry].wasSeen = true
				if !d.rooms[rx][ry].isGenerated {
					d.rooms[rx][ry].generateDungeonCell()
				}
			}
		}
	}
}

func (d *dungeon) canPlayerMoveFromByVector(p *player, vx, vy int) bool {
	elem := d.layout.GetElement(p.dungX, p.dungY)
	conns := elem.GetAllConnectionsCoords()
	for c := 0; c < len(conns); c++ {
		if conns[c][0] == vx && conns[c][1] == vy {
			conn := elem.GetConnectionByCoords(vx, vy)
			if !conn.IsLocked || p.keys[conn.LockNum] {
				return true
			}
		}
	}
	return false
}
