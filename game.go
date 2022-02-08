package main

import "fmt"

func runGame() {
	dung := &dungeon{}
	x, y := dung.initAndGenerate("dungeon_generation_patterns/explore_or_fight.ptn")
	plr := &player{
		dungX: x,
		dungY: y,
	}
	plr.init()
	dung.plr = plr
	dung.startDungeonLoop()
}

func (d *dungeon) startDungeonLoop() {
	for {
		d.performPreTurnCellActions()
		if d.plr.hitpoints <= 0 {
			io.showYNSelect("YOU DIED", []string{"Try again?"})
			return
		}
		io.renderDungeon(d, d.plr)
		key := io.readKey()
		vx, vy := readKeyToVector(key)
		d.movePlayerByVector(vx, vy)
		switch key {
		case "ENTER":
			d.pickUpFromRoom(d.getPlayerRoom())
		case "ESCAPE":
			return
		}
	}
}

func (d *dungeon) performPreTurnCellActions() {
	d.generateAndRevealRoomsAroundPlayer()
	room := d.getPlayerRoom()
	if room.hasKey > 0 {
		d.plr.keys[room.hasKey] = true
		room.hasKey = 0
	}
}

func (d *dungeon) doesPlayerEnterRoom(vx, vy int) bool {
	x, y := d.plr.dungX+vx, d.plr.dungY+vy
	room := d.rooms[x][y]
	room.wasSeen = true
	// enter combat?
	if !room.isCleared() {
		if d.offerCombatToPlayer(room) {
			b := generateBattlefield(room, d.plr)
			b.startCombatLoop()
			d.onCombatEnd(b, room)
			return true
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

func (d *dungeon) onCombatEnd(b *battlefield, room *dungeonCell) {
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
		d.pickUpFromRoom(room)
	} else {

	}
}

func (d *dungeon) pickUpFromRoom(r *dungeonCell) {
	for len(r.treasure) > 0 {
		lines := []string{}
		for _, i := range r.treasure {
			if i.IsConsumable() {
				lines = append(lines, fmt.Sprintf("%dx %s", i.AsConsumable.Amount, i.GetName()))
			} else {
				lines = append(lines, fmt.Sprintf(i.GetName()))
			}
		}
		picked := io.showSelectWindow("Pick up:", lines)
		if picked == -1 {
			break
		}
		d.plr.acquireItem(r.treasure[picked])
		r.treasure = append(r.treasure[:picked], r.treasure[picked+1:]...)
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
