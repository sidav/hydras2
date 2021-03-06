package main

import "hydras2/text_colors"

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
		io.showInfoWindow(
			"WELCOME TO HYDRAS 2",
			"You are a Hydra Hunter! With your yet simple weapons you move into this " +
				"dungeon full of strangeness. \n" +
				text_colors.MakeStringColorTagged("Will you be able to save your kingdom from hydras' vermin?", []string{"YELLOW"}),
			)
		dung.startDungeonLoop()
	}
}
