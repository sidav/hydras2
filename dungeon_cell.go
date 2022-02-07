package main

import "hydras2/entities"

type dungeonCell struct {
	isRoom      bool
	isGenerated bool
	isVisited   bool

	enemies  []*enemy
	treasure []*entities.Item
}

func (dc *dungeonCell) generateDungeonCell() {
	if dc.isRoom {
		numEnemies := rnd.RandInRange(0, 3)
		for i := 0; i < numEnemies; i++ {
			dc.enemies = append(dc.enemies, &enemy{
				enemyType: ENEMY_HYDRA,
				heads:     rnd.RandInRange(1, 5),
			})
		}
		numItems := rnd.RandInRange(0, 3)
		for i := 0; i < numItems; i++ {
			dc.treasure = append(dc.treasure, entities.GenerateRandomItem(rnd))
		}
	} else {
	}
	dc.isGenerated = true
}

func (dc *dungeonCell) isCleared() bool {
	return !dc.isGenerated || len(dc.enemies) == 0
}
