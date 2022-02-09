package main

import "hydras2/entities"

func (b *battlefield) usePlayerConsumable() {
	if b.player.currentConsumable.AsConsumable.Amount <= 0 {
		return
	}
	switch b.player.currentConsumable.AsConsumable.Code {
	case entities.CONSUMABLE_HEAL:
		if b.player.hitpoints == b.player.getMaxHp() {
			return
		}
		healAmount := 3 + b.player.currentConsumable.AsConsumable.Enchantment
		b.player.hitpoints += healAmount
		if b.player.hitpoints > b.player.getMaxHp() {
			b.player.hitpoints = b.player.getMaxHp()
		}
	case entities.CONSUMABLE_RETREAT:
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
