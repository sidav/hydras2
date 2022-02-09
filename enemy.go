package main

import (
	"fmt"
	"hydras2/entities"
	"hydras2/text_colors"
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
	name := e.element.GetName()
	if len(name) > 0 {
		name += " "
	}
	name += fmt.Sprintf("%d-headed hydra", e.heads)
	return text_colors.MakeStringColorTagged(name, e.element.GetColorTags())
}

func generateRandomEnemy(minHeads, maxHeads int) *enemy {
	e := &enemy{
		enemyType: ENEMY_HYDRA,
		heads:     rnd.RandInRange(1, 5),
		element:   &entities.Element{Code: entities.GetWeightedRandomElementCode(rnd)},
	}
	e.headsOnGeneration = e.heads
	return e
}
