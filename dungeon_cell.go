package main

import "hydras2/entities"

type dungeonCell struct {
	isRoom            bool
	contentsGenerated bool

	enemies  []*enemy
	treasure []*entities.Item
	feature  *dungeonRoomFeature
	hasKey   int
}

func (dc *dungeonCell) generateDungeonCellContents(discoveryCount int, boss bool) {
	if dc.isRoom {
		numEnemies := rnd.RandInRange(0, discoveryCount/3+1)
		if numEnemies > 5 {
			numEnemies = 5
		}
		minHeads, maxHeads := discoveryCount/5+1, discoveryCount/3+1
		for i := 0; i < numEnemies && len(dc.enemies) < 5; i++ {
			dc.enemies = append(dc.enemies, generateRandomEnemy(minHeads, maxHeads, rnd.OneChanceFrom(5), false, false))
		}
		if boss {
			bossEnemy := generateRandomEnemy(
				discoveryCount*2,
				discoveryCount*3,
				true,
				true,
				true)
			bossEnemy.isBoss = true
			dc.enemies = append(dc.enemies, bossEnemy)
		}

		// add treasure
		if len(dc.treasure) > 0 { // it may be when re-generating
			dc.treasure = make([]*entities.Item, 0)
		}
		numItems := rnd.RandInRange(9, 9)
		if boss {
			numItems += 3
		}
		for i := 0; i < numItems; i++ {
			dc.treasure = append(dc.treasure, entities.GenerateRandomItem(rnd))
		}
	}
	entities.SortItemsArray(dc.treasure)
	dc.contentsGenerated = true
}

func (dc *dungeonCell) hasFeature(ftype int) bool {
	if dc.feature == nil {
		return false
	}
	return dc.feature.featureCode == ftype
}

func (dc *dungeonCell) isCleared() bool {
	return !dc.contentsGenerated || len(dc.enemies) == 0
}
