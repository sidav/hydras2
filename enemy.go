package main

import (
	"fmt"
	"hydras2/entities"
)

const (
	ENEMY_HYDRA = iota
)

type enemy struct {
	enemyType         int
	heads             int
	element           *entities.Element
	headsOnGeneration int // for resetting rooms to initial state

	// battlefield-only vars:
	x, y          int
	dirx, diry    int
	nextTickToAct int
}

func (e *enemy) getName() string {
	return fmt.Sprintf("%d-headed %s hydra", e.heads, e.element.GetName())
}

func generateRandomEnemy() *enemy {
	e := &enemy{
		enemyType: ENEMY_HYDRA,
		heads:     rnd.RandInRange(1, 5),
		element: &entities.Element{Code: entities.GetWeightedRandomElementCode(rnd)},
	}
	e.headsOnGeneration = e.heads
	return e
}
