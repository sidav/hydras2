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

func (b *battlefield) init(dc *dungeonCell, p *player) {
	const bfSize = 10
	b.tiles = make([][]int, bfSize)
	for i := range b.tiles {
		b.tiles[i] = make([]int, bfSize)
	}

	for i := range dc.enemies {
		b.enemies = append(b.enemies, dc.enemies[i])
		b.enemies[i].x = i*2+1
		b.enemies[i].y = 1
	}
	b.player = p
}

func (b *battlefield) startCombat() {
	for {
		io.renderBattlefield(b)
		key := io.readKey()
		switch key {
		case "ESCAPE":
			return

		}
	}
}
