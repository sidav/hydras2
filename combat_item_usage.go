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
		b.player.hitpoints += b.player.currentConsumable.AsConsumable.EnchantAmount
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
	case entities.CONSUMABLE_DESTROY_HYDRA:
		index := rnd.SelectRandomIndexFromWeighted(len(b.enemies), func(x int) int {return b.enemies[x].heads})
		hydraToDestroy := b.enemies[index]
		if hydraToDestroy.aura != nil {
			log.AppendMessagef("Powerful %s is only half-decapitated!", hydraToDestroy)
			hydraToDestroy.heads /= 2
		} else {
			log.AppendMessagef("%s is obliterated!!", hydraToDestroy)
			hydraToDestroy.heads = 0
		}
	case entities.CONSUMABLE_SHUFFLE_ELEMENT:
		for _, e := range b.enemies {
			e.element.Code = entities.GetWeightedRandomElementCode(rnd)
		}
		log.AppendMessagef("Hydras have their elements shifted!")
	}

	log.AppendMessagef("%s used.", b.player.currentConsumable.GetName())
	b.player.currentConsumable.AsConsumable.Amount--
}
