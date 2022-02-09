package main

import (
	"github.com/sidav/cyclicdungeongenerator/generator"
)

type dungeon struct {
	plr              *player
	startX, startY   int
	totalOpenedRooms int
	name             string
	layout           generator.LayoutInterface
	rooms            [][]*dungeonCell
}

func (d *dungeon) initAndGenerate(patternFileName string) {
	const sizew, sizeh = 7, 4
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
				d.rooms[x][y].contentsGenerated = true
				d.rooms[x][y].feature = &dungeonRoomFeature{featureCode: DRF_BONFIRE}
				d.startX = x
				d.startY = y
			}
			if d.layout.GetElement(x, y).IsNode() && d.layout.GetElement(x, y).HasTag("ky1") {
				d.rooms[x][y].hasKey = 1
			}
			if d.layout.GetElement(x, y).IsNode() && d.layout.GetElement(x, y).HasTag("ky2") {
				d.rooms[x][y].hasKey = 2
			}
		}
	}
	d.placeFeatureInRandomRoom(DRF_ALTAR)
	d.placeFeatureInRandomRoom(DRF_TINKER)
}

func (d *dungeon) placeFeatureInRandomRoom(featureCode int) {
	for try := 0; try < 1000; try++ {
		x, y := rnd.Rand(len(d.rooms)), rnd.Rand(len(d.rooms[0]))
		if !d.rooms[x][y].isRoom || d.rooms[x][y].feature != nil {
			continue
		}
		d.rooms[x][y].feature = &dungeonRoomFeature{
			featureCode: featureCode,
		}
		return
	}
	panic("Feature can't be placed!")
}

func (d *dungeon) getPlayerRoom() *dungeonCell {
	return d.rooms[d.plr.dungX][d.plr.dungY]
}

func (d *dungeon) generateAndRevealRoomsAroundPlayer() {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if (x == 0 || y == 0) && d.canPlayerMoveFromByVector(x, y) {
				rx, ry := d.plr.dungX+x, d.plr.dungY+y
				d.rooms[rx][ry].wasSeen = true
				if !d.rooms[rx][ry].contentsGenerated {
					if d.rooms[rx][ry].isRoom {
						d.totalOpenedRooms++
					}
					d.rooms[rx][ry].generateDungeonCell(d.totalOpenedRooms)
				}
			}
		}
	}
}

func (d *dungeon) clearRoomsGeneratedState() {
	for x := range d.rooms {
		for y := range d.rooms[x] {
			if d.rooms[x][y].isCleared() && !d.rooms[x][y].hasFeature(DRF_BONFIRE) {
				d.rooms[x][y].contentsGenerated = false
				d.rooms[x][y].wasSeen = false
			}
		}
	}
}

func (d *dungeon) canPlayerMoveFromByVector(vx, vy int) bool {
	if vx == 0 && vy == 0 {
		return true
	}
	elem := d.layout.GetElement(d.plr.dungX, d.plr.dungY)
	conns := elem.GetAllConnectionsCoords()
	for c := 0; c < len(conns); c++ {
		if conns[c][0] == vx && conns[c][1] == vy {
			conn := elem.GetConnectionByCoords(vx, vy)
			if !conn.IsLocked || d.plr.keys[conn.LockNum] {
				return true
			}
		}
	}
	return false
}
