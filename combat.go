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
}

func generateBattlefield(dc *dungeonCell, p *player) *battlefield {
	b := &battlefield{}
	bfW := rnd.RandInRange(3, 5) * 2 + 1
	bfH := rnd.RandInRange(5, 9)
	b.tiles = make([][]int, bfW)
	for i := range b.tiles {
		b.tiles[i] = make([]int, bfH)
	}

	for i := range dc.enemies {
		b.enemies = append(b.enemies, dc.enemies[i])
		b.enemies[i].nextTickToAct = 0
		b.enemies[i].x = i*2+1
		b.enemies[i].y = 1
	}
	b.player = p
	b.player.x = bfW/2
	b.player.y = bfH-2
	return b
}

func (b *battlefield) startCombatLoop() {
	for {
		io.renderBattlefield(b)
		key := io.readKey()
		switch key {
		case "ESCAPE":
			return

		}
	}
}
