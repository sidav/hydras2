package main

import "hydras2/entities"

type dungeonCell struct {
	isRoom            bool
	contentsGenerated bool
	wasSeen           bool

	enemies  []*enemy
	treasure []*entities.Item
	feature  *dungeonRoomFeature
	hasKey   int
}

func (dc *dungeonCell) generateDungeonCell(num int) {
	if dc.isRoom {
		numEnemies := rnd.RandInRange(0, num/3+1)
		if numEnemies > 5 {
			numEnemies = 5
		}
		minHeads, maxHeads := num/5+1, num/3+1
		for i := 0; i < numEnemies && len(dc.enemies) < 5; i++ {
			dc.enemies = append(dc.enemies, generateRandomEnemy(minHeads, maxHeads, rnd.OneChanceFrom(5), false))
		}
		if len(dc.treasure) > 0 { // it may be when re-generating
			return
		} else {
			numItems := rnd.RandInRange(0, 3)
			for i := 0; i < numItems; i++ {
				dc.treasure = append(dc.treasure, entities.GenerateRandomItem(rnd))
			}
		}
	} else {
	}
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
