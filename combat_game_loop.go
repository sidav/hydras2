package main

func (b *battlefield) startCombatLoop() {
	log.Clear()
	for !b.battleEnded {
		for b.player.nextTickToAct <= b.currentTick && !b.battleEnded {
			io.renderBattlefield(b)
			b.workPlayerInput()
		}
		b.cleanup()
		for _, e := range b.enemies {
			if e.nextTickToAct <= b.currentTick {
				b.actAsEnemy(e)
			}
		}
		b.activatePassiveBrandsOnPlayerItems()
		b.currentTick++
	}
	log.Clear()
}

func (b *battlefield) workPlayerInput() {
	correctInputKeyPressed := false
	for !correctInputKeyPressed {
		key := io.readKey()
		switch key {
		case "ESCAPE":
			b.battleEnded = true
			return
		case "1":
			b.player.cycleToNextPrimaryWeapon()
			return
		case "2":
			b.player.cycleToNextSecondaryWeapon()
			return
		case "3":
			b.player.cycleToNextConsumable()
			return
		case "x":
			b.player.swapWeapons()
			return
		case "ENTER":
			b.usePlayerConsumable()
			b.player.nextTickToAct = b.currentTick + 10
			return
		case " ":
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
							b.playerHitsEnemy(enemyAt)
							b.player.nextTickToAct = b.currentTick + 10
						} else {
							b.player.x += vx
							b.player.y += vy
							b.player.nextTickToAct = b.currentTick + 10
						}
					}
				}
			}
		}
	}
}

func (b *battlefield) playerHitsEnemy(e *enemy) {
	weaponToAttackWith := b.player.primaryWeapon
	if b.player.primaryWeapon.AsWeapon.GetDamageOnHeads(e.heads) == 0 {
		weaponToAttackWith = b.player.secondaryWeapon
	}
	b.performWeaponStrikeOnEnemy(weaponToAttackWith.AsWeapon, e)
	b.activateOnHitBrandOnItem(weaponToAttackWith, e)
}
