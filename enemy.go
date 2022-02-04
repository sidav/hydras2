package main

import "fmt"

const (
	ENEMY_HYDRA = iota
)

type enemy struct {
	enemyType int
	heads     int
}

func (e *enemy) getName() string {
	return fmt.Sprintf("%d-headed hydra", e.heads)
}
