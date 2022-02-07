package main

var (
	dung *dungeon
	plr  *player
)

func runGame() {
	dung = &dungeon{}
	x, y := dung.initAndGenerate("dungeon_generation_patterns/explore_or_fight.ptn")
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
		if plr.hitpoints <= 0 {
			io.showYNSelect("YOU DIED", []string{"Try again?"})
			return
		}
		io.renderDungeon(dung, plr)
		key := io.readKey()
		vx, vy := readKeyToVector(key)
		movePlayerByVector(vx, vy)
		switch key {
		case "ESCAPE":
			return
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
	// enter combat?
	if !dung.rooms[x][y].isCleared() {
		var lines []string
		lines = append(lines, "   Enemies:")
		for _, e := range dung.rooms[x][y].enemies {
			lines = append(lines, e.getName())
		}
		if len(dung.rooms[x][y].treasure) > 0 {
			lines = append(lines, "   Treasure:")
			for _, t := range dung.rooms[x][y].treasure {
				lines = append(lines, t.GetName())
			}
		}
		lines = append(lines, "")
		lines = append(lines, "  Enter the combat?")
		if io.showYNSelect(" ENCOUNTER ", lines) {
			b := generateBattlefield(dung.rooms[x][y], plr)
			b.startCombatLoop()
			// clear room enemies if player defeated them
			if len(b.enemies) == 0 {
				dung.rooms[x][y].enemies = []*enemy{}
			}
			return true
		}
		return false
	}
	return true
}

func onCombatEnd() {

}

func movePlayerByVector(vx, vy int) {
	if dung.canPlayerMoveFromByVector(plr, vx, vy) {
		if onCellEntry(vx, vy) {
			plr.dungX += vx
			plr.dungY += vy
		}
	}
}
