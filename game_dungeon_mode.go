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
	dung.rooms[plr.dungX][plr.dungY].isVisited = true
}

func onCellEntry(vx, vy int) bool {
	x, y := plr.dungX+vx, plr.dungY+vy
	if !dung.rooms[x][y].isGenerated {
		dung.rooms[x][y].generateDungeonCell()
	}
	dung.rooms[x][y].isVisited = true
	if !dung.rooms[x][y].isCleared() {
		var lines []string
		for _, e := range dung.rooms[x][y].enemies {
			lines = append(lines, e.getName())
		}
		lines = append(lines, "   Treasure:")
		for _, t := range dung.rooms[x][y].treasure {
			lines = append(lines, t.getName())
		}
		lines = append(lines, "   Enter the combat?")
		return r.showYNSelect("  You see here enemies:", lines)
	}
	return true
}

func movePlayerByVector(vx, vy int) {
	if dung.canPlayerMoveFromByVector(plr, vx, vy) {
		if onCellEntry(vx, vy) {
			plr.dungX += vx
			plr.dungY += vy
		}
	}
}
