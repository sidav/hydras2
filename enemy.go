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
	isBoss            bool
	aura              *entities.Aura
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
	auraName := ""
	if e.aura != nil {
		auraName = e.aura.GetName() + " "
	}
	name += fmt.Sprintf("%d-headed %shydra", e.heads, auraName)
	if e.isBoss {
		name += " overlord"
	}
	return text_colors.MakeStringColorTagged(name, e.element.GetColorTags())
}

func generateRandomEnemy(minHeads, maxHeads int, aura, prime bool) *enemy {
	heads := rnd.RandInRange(minHeads, maxHeads)
	if prime {
		heads = rnd.GenerateRandomPrimeInRange(minHeads, maxHeads)
	}
	e := &enemy{
		enemyType: ENEMY_HYDRA,
		heads:     heads,
		element:   &entities.Element{Code: entities.GetWeightedRandomElementCode(rnd)},
	}
	if aura {
		e.aura = entities.GenerateRandomAura(rnd)
	}
	e.headsOnGeneration = e.heads
	return e
}
