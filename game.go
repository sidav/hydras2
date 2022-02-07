package main

import "fmt"

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
	room := dung.rooms[plr.dungX][plr.dungY]
	room.isVisited = true
	if room.hasKey > 0 {
		plr.keys[room.hasKey] = true
		room.hasKey = 0
	}
}

func onCellEntry(vx, vy int) bool {
	x, y := plr.dungX+vx, plr.dungY+vy
	room := dung.rooms[x][y]
	if !room.isGenerated {
		room.generateDungeonCell()
	}
	room.isVisited = true
	// enter combat?
	if !room.isCleared() {
		if offerCombatToPlayer(room) {
			b := generateBattlefield(room, plr)
			b.startCombatLoop()
			onCombatEnd(b, room)
			return true
		}
		return false
	}
	return true
}

func offerCombatToPlayer(room *dungeonCell) bool {
	var lines []string
	lines = append(lines, "   Enemies:")
	for _, e := range room.enemies {
		lines = append(lines, e.getName())
	}
	if len(room.treasure) > 0 {
		lines = append(lines, "   Treasure:")
		for _, t := range room.treasure {
			lines = append(lines, t.GetName())
		}
	}
	if room.hasKey > 0 {
		lines = append(lines, " !There is a key!")
	}
	lines = append(lines, "")
	lines = append(lines, "  Enter the combat?")
	return io.showYNSelect(" ENCOUNTER ", lines)
}

func onCombatEnd(b *battlefield, room *dungeonCell) {
	if len(b.enemies) == 0 {
		soulsAcquired := 0
		for i := range room.enemies {
			soulsAcquired += room.enemies[i].headsOnGeneration
		}
		lines := []string {
			fmt.Sprintf("%d hydras slain.", len(room.enemies)),
			fmt.Sprintf("You acquired %d hydra essense", soulsAcquired),
		}
		if room.hasKey > 0 {
			lines = append(lines, fmt.Sprintf("Acquired key %d", room.hasKey))
		}
		io.showInfoWindow("VICTORY ACHIEVED", lines)
		room.enemies = []*enemy{}
	} else {

	}
}

func movePlayerByVector(vx, vy int) {
	if dung.canPlayerMoveFromByVector(plr, vx, vy) {
		if onCellEntry(vx, vy) {
			plr.dungX += vx
			plr.dungY += vy
		}
	}
}
