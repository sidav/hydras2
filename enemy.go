package main

import "fmt"

const (
	ENEMY_HYDRA = iota
)

type enemy struct {
	enemyType         int
	heads             int
	headsOnGeneration int // for resetting rooms to initial state

	// battlefield-only vars:
	x, y          int
	dirx, diry    int
	nextTickToAct int
}

func (e *enemy) getName() string {
	return fmt.Sprintf("%d-headed hydra", e.heads)
}

func generateRandomEnemy() *enemy {
	e := &enemy{
		enemyType: ENEMY_HYDRA,
		heads:     rnd.RandInRange(1, 5),
	}
	e.headsOnGeneration = e.heads
	return e
}
