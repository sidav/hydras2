package main

type dungeonCell struct {
	isRoom      bool
	isCleared   bool
	isGenerated bool
	isVisited   bool

	enemies  []*enemy
	treasure []*item
}

func (dc *dungeonCell) generateDungeonCell() {
	numEnemies := rnd.RandInRange(0, 3)
	for i := 0; i < numEnemies; i++ {
		dc.enemies = append(dc.enemies, &enemy{
			enemyType: ENEMY_HYDRA,
			heads:     rnd.RandInRange(1, 5),
		})
	}
	numItems := rnd.RandInRange(0, 3)
	for i := 0; i < numItems; i++ {
		dc.treasure = append(dc.treasure, generateRandomItem())
	}
}
