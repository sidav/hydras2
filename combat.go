package main

import (
	"hydras2/entities"
	"hydras2/text_colors"
	"math"
)

const (
	TILE_FLOOR = iota
	TILE_WALL
)

type battlefield struct {
	tiles       [][]int
	enemies     []*enemy
	player      *player
	currentTick int

	battleEnded bool
	playerFled  bool
}

func generateBattlefield(p *player, enemies []*enemy) *battlefield {
	b := &battlefield{}
	bfW := rnd.RandInRange(3, 5)*2 + 1
	bfH := rnd.RandInRange(5, 9)
	b.tiles = make([][]int, bfW)
	for i := range b.tiles {
		b.tiles[i] = make([]int, bfH)
	}

	for i := 0; i < 20*bfW*bfH/100; i++ {
		x, y := rnd.RandInRange(1, bfW-2), rnd.RandInRange(1, bfH-2)
		b.tiles[x][y] = TILE_WALL
	}

	for i := range enemies {
		b.enemies = append(b.enemies, enemies[i])
		b.enemies[i].nextTickToAct = 0
		b.enemies[i].x = i*2 + 1
		b.enemies[i].y = 0
	}
	b.player = p
	b.player.x = bfW / 2
	b.player.y = bfH - 1
	b.player.nextTickToAct = 1
	return b
}

func (b *battlefield) getEnemyAt(x, y int) *enemy {
	for i := range b.enemies {
		if b.enemies[i].x == x && b.enemies[i].y == y {
			return b.enemies[i]
		}
	}
	return nil
}

func (b *battlefield) actAsEnemy(e *enemy) {
	const faceChangePeriod = 10
	// first, check if we're in same row or col with the player
	if e.x == b.player.x || e.y == b.player.y {
		// if so, set direction to player
		e.dirx, e.diry = vectorToUnitVector(b.player.x-e.x, b.player.y-e.y)
	}
	newX, newY := e.x+e.dirx, e.y+e.diry
	// if random or if the cell is unpassable, rotate randomly
	if !b.areCoordsValid(newX, newY) || b.tiles[newX][newY] == TILE_WALL || rnd.OneChanceFrom(faceChangePeriod) ||
		e.dirx == 0 && e.diry == 0 || b.getEnemyAt(newX, newY) != nil {
		e.dirx, e.diry = rnd.RandomUnitVectorInt(false)
	}
	newX, newY = e.x+e.dirx, e.y+e.diry
	if b.areCoordsValid(newX, newY) && b.tiles[newX][newY] != TILE_WALL && b.getEnemyAt(newX, newY) == nil {
		if b.player.x == newX && b.player.y == newY {
			b.enemyHitsPlayer(e)
			e.nextTickToAct = b.currentTick + 10
		} else {
			e.x += e.dirx
			e.y += e.diry
			e.nextTickToAct = b.currentTick + 10
		}
	}
}

func (b *battlefield) areCoordsValid(x, y int) bool {
	return x >= 0 && x < len(b.tiles) && y >= 0 && y < len(b.tiles[0])
}

func (b *battlefield) performWeaponStrikeOnEnemy(weapon *entities.ItemWeapon, e *enemy) {
	dmg := weapon.GetDamageOnHeads(e.heads)
	log.AppendMessagef("You cut %d heads off %s!", dmg, e.getName())
	if e.heads <= dmg {
		log.AppendMessagef("%s drops dead!", e.getName())
	}
	e.heads -= dmg
	playerWeaponElement := weapon.WeaponElement
	hydraElement := e.element
	if e.heads > 0 {
		switch playerWeaponElement.GetEffectivenessAgainstElement(hydraElement) {
		case -1:
			log.AppendMessagef("%s doubles its heads!", e.getName())
			e.heads *= 2
		case 0:
			log.AppendMessagef("%s grows a head!", e.getName())
			e.heads++
		case 1:
			log.AppendMessagef("%s writhes!", e.getName())
		}
	}
}

func (b *battlefield) enemyHitsPlayer(e *enemy) {
	dmg := int(math.Log2(float64(e.heads + 1)))
	b.player.hitpoints -= dmg
	log.AppendMessagef("%s bites you for %d damage!", e.getName(), dmg)
	if b.player.hitpoints <= b.player.getMaxHp()/3 {
		log.AppendMessage(text_colors.MakeStringColorTagged("!!LOW HITPOINT WARNING!!", []string{"RED"}))
	}
}

func (b *battlefield) cleanup() {
	for i := 0; i < len(b.enemies); i++ {
		if b.enemies[i].heads <= 0 {
			b.enemies[i] = b.enemies[len(b.enemies)-1]
			b.enemies = b.enemies[:len(b.enemies)-1]
			i--
		}
	}
	if b.player.hitpoints <= 0 || len(b.enemies) == 0 {
		b.battleEnded = true
	}
}
