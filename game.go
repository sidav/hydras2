package main

func runGame() {
	for !exitGame {
		dung := &dungeon{}
		dung.initAndGenerate("dungeon_generation_patterns/explore_or_fight.ptn")
		plr := &player{
			dungX: dung.startX,
			dungY: dung.startY,
		}
		plr.init()
		dung.plr = plr
		dung.startDungeonLoop()
	}
}
