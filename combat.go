package main

const (
	TILE_FLOOR = iota
	TILE_WALL
)

type battlefield struct {
	tiles [][]int
	enemies []*enemy
	player *player
	currentTick int

	battleEnded bool
}

func generateBattlefield(dc *dungeonCell, p *player) *battlefield {
	b := &battlefield{}
	bfW := rnd.RandInRange(3, 5) * 2 + 1
	bfH := rnd.RandInRange(5, 9)
	b.tiles = make([][]int, bfW)
	for i := range b.tiles {
		b.tiles[i] = make([]int, bfH)
	}

	for i := 0; i < 20*bfW*bfH/100; i++ {
		x, y := rnd.RandInRange(1, bfW-2), rnd.RandInRange(1, bfH-2)
		b.tiles[x][y] = TILE_WALL
	}

	for i := range dc.enemies {
		b.enemies = append(b.enemies, dc.enemies[i])
		b.enemies[i].nextTickToAct = 0
		b.enemies[i].x = i*2+1
		b.enemies[i].y = 0
	}
	b.player = p
	b.player.x = bfW/2
	b.player.y = bfH-2
	return b
}

func (b *battlefield) startCombatLoop() {
	for !b.battleEnded {
		io.renderBattlefield(b)
		b.workPlayerInput()
		for _, e := range b.enemies {
			b.actAsEnemy(e)
		}
		b.currentTick++
	}
}

func (b *battlefield) getEnemyAt(x, y int) *enemy {
	for i := range b.enemies {
		if b.enemies[i].x == x && b.enemies[i].y == y {
			return b.enemies[i]
		}
	}
	return nil
}

func (b *battlefield) workPlayerInput() {
	correctInputKeyPressed := false
	for !correctInputKeyPressed {
		key := io.readKey()
		switch key {
		case "ESCAPE":
			correctInputKeyPressed = true
			b.battleEnded = true
			return
		default:
			vx, vy := readKeyToVector(key)
			if vx != 0 || vy != 0 {
				correctInputKeyPressed = true
				newPlrX, newPlrY := b.player.x+vx, b.player.y+vy
				if b.areCoordsValid(newPlrX, newPlrY) {
					if b.tiles[newPlrX][newPlrY] != TILE_WALL {
						enemyAt := b.getEnemyAt(newPlrX, newPlrY)
						if enemyAt != nil {
							return
						}
						b.player.x += vx
						b.player.y += vy
					}
				}
			}
		}
	}
}

func (b *battlefield) actAsEnemy(e *enemy) {
	const faceChangePeriod = 10
	// first, check if we're in same row or col with the player
	if e.x == b.player.x || e.y == b.player.y {
		// if so, set direction to player
		e.dirx, e.diry = vectorToUnitVector(b.player.x-e.x, b.player.y-e.y)
	}
	newX, newY := e.x + e.dirx, e.y+e.diry
	// if random or if the cell is unpassable, rotate randomly
	if !b.areCoordsValid(newX, newY) || b.tiles[newX][newY] == TILE_WALL || rnd.OneChanceFrom(faceChangePeriod) || e.dirx == 0 && e.diry == 0 {
		e.dirx, e.diry = rnd.RandomUnitVectorInt(false)
	}
	newX, newY = e.x + e.dirx, e.y+e.diry
	if b.areCoordsValid(newX, newY) && b.tiles[newX][newY] != TILE_WALL {
		if b.player.x == newX && b.player.y == newY {

		} else {
			e.x += e.dirx
			e.y += e.diry
		}
	}
}

func (b *battlefield) areCoordsValid(x, y int) bool {
	return x >= 0 && x < len(b.tiles) && y >= 0 && y < len(b.tiles[0])
}
