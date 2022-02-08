package main

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
