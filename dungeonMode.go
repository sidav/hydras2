package main

var (
	dung *dungeon
	plr *player
)

func runGame() {
	dung = &dungeon{}
	x, y := dung.initAndGenerate("patterns/explore_or_fight.ptn")
	plr = &player{
		dungX: x,
		dungY: y,
	}
	dungeonMode()
}

func dungeonMode() {
	for {
		r.renderDungeon(dung, plr)
		key := r.readKey()
		switch key {
		case "ESCAPE":
			return
		case "UP":
			if dung.canPlayerMoveFromByVector(plr, 0, -1) {
				plr.dungY--
			}
		case "DOWN":
			if dung.canPlayerMoveFromByVector(plr, 0, 1) {
				plr.dungY++
			}
		case "LEFT":
			if dung.canPlayerMoveFromByVector(plr, -1, 0) {
				plr.dungX--
			}
		case "RIGHT":
			if dung.canPlayerMoveFromByVector(plr, 1, 0) {
				plr.dungX++
			}
		}
	}
}
