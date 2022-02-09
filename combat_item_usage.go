package main

import "hydras2/entities"

func (b *battlefield) usePlayerConsumable() {
	if b.player.currentConsumable.AsConsumable.Amount <= 0 {
		return
	}
	switch b.player.currentConsumable.AsConsumable.Code {
	case entities.ITEM_HEAL:
		if b.player.hitpoints == b.player.getMaxHp() {
			return
		}
		b.player.hitpoints += 4
		if b.player.hitpoints > b.player.getMaxHp() {
			b.player.hitpoints = b.player.getMaxHp()
		}
	case entities.ITEM_RETREAT:
		if io.showYNSelect("Flee?", []string{"You will lose all essence!"}) {
			b.playerFled = true
			b.battleEnded = true
		} else {
			return
		}
	}

	log.AppendMessagef("%s used.", b.player.currentConsumable.GetName())
	b.player.currentConsumable.AsConsumable.Amount--
}
