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
		case entities.BRAND_PASSIVE_DISTORTION:
			item.AsWeapon.Damage += rnd.RandInRange(-1, 1)
			if item.AsWeapon.Damage == 0 {
				item.AsWeapon.Damage = 1
			}
		}
	}
}
