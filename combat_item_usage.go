package main

import "hydras2/entities"

func (b *battlefield) usePlayerConsumable() {
	if b.player.currentConsumable.AsConsumable.Amount <= 0 {
		return
	}
	switch b.player.currentConsumable.AsConsumable.Code {
	case entities.ITEM_HEAL:
		b.player.hitpoints += 5
		b.player.hitpoints = b.player.hitpoints % b.player.getMaxHp()
	}
	b.player.currentConsumable.AsConsumable.Amount--
}
