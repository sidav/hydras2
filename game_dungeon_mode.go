package main

var (
	dung *dungeon
	plr  *player
)

func runGame() {
	dung = &dungeon{}
	x, y := dung.initAndGenerate("patterns/explore_or_fight.ptn")
	plr = &player{
		dungX: x,
		dungY: y,
	}
	plr.init()
	dungeonMode()
}

func dungeonMode() {
	for {
		performCellActions()
		r.renderDungeon(dung, plr)
		key := r.readKey()
		switch key {
		case "ESCAPE":
			return
		case "UP":
			movePlayerByVector(0, -1)
		case "DOWN":
			movePlayerByVector(0, 1)
		case "LEFT":
			movePlayerByVector(-1, 0)
		case "RIGHT":
			movePlayerByVector(1, 0)
		}
	}
}

func performCellActions() {
	if !dung.rooms[plr.dungX][plr.dungY].isGenerated {
		dung.rooms[plr.dungX][plr.dungY].generateDungeonCell()
	}
	dung.rooms[plr.dungX][plr.dungY].isVisited = true
}

func movePlayerByVector(vx, vy int) {
	if dung.canPlayerMoveFromByVector(plr, vx, vy) {
		plr.dungX += vx
		plr.dungY += vy
	}
}
