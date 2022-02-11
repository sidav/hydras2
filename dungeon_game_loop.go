package main

import "fmt"

func (d *dungeon) startDungeonLoop() {
	for {
		d.performPreTurnCellActions()
		if d.plr.hitpoints <= 0 {
			exitGame = !io.showYNSelect("YOU DIED", []string{"Try again?"})
			return
		}
		io.renderDungeon(d, d.plr)
		key := io.readKey()
		switch key {
		case "ENTER":
			d.selectPlayerRoomAction()
		case "ESCAPE":
			exitGame = true
			return
		}
		vx, vy := readKeyToVector(key)
		d.movePlayerByVector(vx, vy)
	}
}

func (d *dungeon) performPreTurnCellActions() {
	log.Clear()
	d.generateAndRevealRoomsAroundPlayer()
	room := d.getPlayerRoom()
	if room.hasKey > 0 {
		log.AppendMessagef("You picked up key %d", room.hasKey)
		d.plr.keys[room.hasKey] = true
		room.hasKey = 0
	}
	if room.feature != nil {
		log.AppendMessagef("There is %s here.", getDungeonRoomFeatureNameByCode(room.feature.featureCode))
	}
}

func (d *dungeon) doesPlayerEnterRoom(vx, vy int) bool {
	x, y := d.plr.dungX+vx, d.plr.dungY+vy
	room := d.rooms[x][y]
	// enter combat?
	if !room.isCleared() {
		if d.offerCombatToPlayer(room) {
			b := generateBattlefield(d.plr, room.enemies)
			b.startCombatLoop()
			return d.onCombatEnd(b, room)
		}
		return false
	}
	return true
}

func (d *dungeon) offerCombatToPlayer(room *dungeonCell) bool {
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

func (d *dungeon) onCombatEnd(b *battlefield, room *dungeonCell) bool {
	if len(b.enemies) == 0 {
		soulsAcquired := 0
		for i := range room.enemies {
			soulsAcquired += room.enemies[i].headsOnGeneration
		}
		d.plr.souls += soulsAcquired
		lines := []string{
			fmt.Sprintf("%d hydras slain.", len(room.enemies)),
			fmt.Sprintf("You acquired %d hydra essense", soulsAcquired),
		}
		if room.hasKey > 0 {
			lines = append(lines, fmt.Sprintf("Acquired key %d", room.hasKey))
		}
		io.showInfoWindow("VICTORY ACHIEVED", lines)
		room.enemies = []*enemy{}
		d.checkGameWon()
		return true
	} else {
		if b.playerFled {
			io.showInfoWindow("YOU HAVE FLED", []string{fmt.Sprintf("You have lost %d essence.", d.plr.souls)})
			d.plr.dungX, d.plr.dungY = d.startX, d.startY
			d.plr.souls = 0
			for i := range room.enemies {
				room.enemies[i].heads = room.enemies[i].headsOnGeneration
			}
			return false
		}
		return true // todo: remove this cheat
	}
}

func (d *dungeon) movePlayerByVector(vx, vy int) {
	if d.canPlayerMoveFromByVector(vx, vy) {
		if d.doesPlayerEnterRoom(vx, vy) {
			d.plr.dungX += vx
			d.plr.dungY += vy
		}
	}
}

func (d *dungeon) checkGameWon() {
	for x := range d.rooms {
		for y := range d.rooms[x] {
			if d.rooms[x][y].isRoom && (!d.rooms[x][y].contentsGenerated || !d.rooms[x][y].isCleared()) {
				return
			}
		}
	}
	exitGame = io.showYNSelect("YOU HAVE WON!",
		[]string{
			"You have saved the kingdom from",
			"the hydra vermin. You now can",
			"exit game or stay here and rest",
			"at bonfire for new and harder",
			"enemies spawn. ",
			"",
			"Do you want to stay here? ",
		},
	)
}
