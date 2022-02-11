package main

import "hydras2/entities"

func (b *battlefield) activatePassiveBrandsOnPlayerItems() {
	b.activatePassiveBrandOnItem(b.player.primaryWeapon)
	b.activatePassiveBrandOnItem(b.player.secondaryWeapon)
}

func (b *battlefield) activateOnHitBrandOnItem(item *entities.Item, hitEnemy *enemy) {
	if item.IsWeapon() {
		if item.AsWeapon.Brand == nil {
			return
		}
		switch item.AsWeapon.Brand.Code {
		case entities.BRAND_DOUBLE_STRIKE:
			log.AppendMessage("The weapon hits twice!")
			b.performWeaponStrikeOnEnemy(item.AsWeapon, hitEnemy)
		case entities.BRAND_SAFE_DOUBLE_STRIKE:
			if item.AsWeapon.GetDamageOnHeads(hitEnemy.heads) > 0 {
				log.AppendMessage("The weapon hits twice!")
				b.performWeaponStrikeOnEnemy(item.AsWeapon, hitEnemy)
			}
		case entities.BRAND_FAST_STRIKE:
			b.player.nextTickToAct -= item.AsWeapon.Brand.EnchantAmount
		case entities.BRAND_VAMPIRISM:
			if rnd.OneChanceFrom(10) && b.player.getMaxHp() > b.player.hitpoints {
				b.player.hitpoints += 1
				log.AppendMessage("You consume hydra's blood!")
			}
		case entities.BRAND_BETTER_VAMPIRISM:
			if rnd.OneChanceFrom(10) {
				b.player.hitpoints += 2
				log.AppendMessage("You devour hydra's blood!!")
				if b.player.hitpoints > b.player.getMaxHp() {
					b.player.hitpoints = b.player.getMaxHp()
				}
			}
		}
	}
}

func (b *battlefield) activatePassiveBrandOnItem(item *entities.Item) {
	if item.IsWeapon() {
		if item.AsWeapon.Brand == nil {
			return
		}
		switch item.AsWeapon.Brand.Code {
		case entities.BRAND_PASSIVE_ELEMENTS_SHIFTING:
			item.AsWeapon.WeaponElement.Code = entities.GetRandomElementCode(rnd)
		case entities.BRAND_PASSIVE_RARE_ELEMENTS_SHIFTING:
			item.AsWeapon.WeaponElement.Code = entities.GetRandomRareElementCode(rnd)
		case entities.BRAND_PASSIVE_DISTORTION:
			item.AsWeapon.Damage += rnd.RandInRange(-1, 1)
			if item.AsWeapon.Damage == 0 {
				item.AsWeapon.Damage = 1
			}
		case entities.BRAND_PASSIVE_SMARTER_DISTORTION:
			highestHeads, lowestHeads := 0, 999
			for _, e := range b.enemies {
				if e.heads < lowestHeads {
					lowestHeads = e.heads
				}
				if e.heads > highestHeads {
					highestHeads = e.heads
				}
			}
			item.AsWeapon.Damage = rnd.RandInRange(1, highestHeads)
			if item.AsWeapon.WeaponTypeCode != entities.WTYPE_DIVISOR {
				item.AsWeapon.Damage = rnd.RandInRange(1, highestHeads/2)
			}
			if item.AsWeapon.Damage == 0 {
				item.AsWeapon.Damage = 1
			}
		}
	}
}
