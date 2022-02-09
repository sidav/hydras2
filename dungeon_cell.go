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

func (dc *dungeonCell) generateDungeonCell() {
	if dc.isRoom {
		numEnemies := rnd.RandInRange(0, 3)
		for i := 0; i < numEnemies; i++ {
			dc.enemies = append(dc.enemies, generateRandomEnemy())
		}
		numItems := rnd.RandInRange(0, numEnemies+1)
		for i := 0; i < numItems; i++ {
			dc.treasure = append(dc.treasure, entities.GenerateRandomItem(rnd))
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
