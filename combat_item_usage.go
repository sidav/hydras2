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
		if io.showYNSelect("Flee?", "Are you sure you want to flee?\n You will lose all essence!") {
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
			e.element.Code = entities.GetRandomElementCode(rnd)
		}
		log.AppendMessagef("Hydras have their elements shifted!")
	case entities.CONSUMABLE_MASS_BISECTION:
		for _, e := range b.enemies {
			e.heads /= 2
		}
		log.AppendMessagef("All hydras are bisected!")
	case entities.CONSUMABLE_PARALYZE_HYDRAS:
		for _, e := range b.enemies {
			e.nextTickToAct = b.currentTick+5*COMBAT_MOVE_COST
		}
		log.AppendMessagef("All hydras are paralyzed for 5 turns!")
	case entities.CONSUMABLE_UNELEMENT_HYDRAS:
		for _, e := range b.enemies {
			e.element.Code = entities.ELEMENT_NONE
		}
		log.AppendMessagef("All hydras are purged!")
	case entities.CONSUMABLE_MERGE_HYDRAS:
		totalHeads := 0
		var enemyToMergeInto *enemy
		for _, e := range b.enemies {
			totalHeads += e.heads
			if enemyToMergeInto == nil || e.aura != nil || e.heads > enemyToMergeInto.heads {
				enemyToMergeInto = e
			}
			e.heads = 0
		}
		enemyToMergeInto.heads = totalHeads
		log.AppendMessagef("All hydras are merged!")

	default:
		panic("wat")
	}

	log.AppendMessagef("%s used.", b.player.currentConsumable.GetName())
	b.player.removeItemFromInventory(b.player.currentConsumable)
}
