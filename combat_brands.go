package main

import "hydras2/entities"

func (b *battlefield) activateBrandsOnPlayerItems() {
	b.activateBrandOnItem(b.player.primaryWeapon)
	b.activateBrandOnItem(b.player.secondaryWeapon)
}

func (b *battlefield) activateBrandOnItem(item *entities.Item) {
	if item.IsWeapon() {
		if item.AsWeapon.Brand == nil {
			return
		}
		switch item.AsWeapon.Brand.Code {
		case entities.BRAND_PASSIVE_ELEMENTS_SHIFTING:
			item.AsWeapon.WeaponElement.Code = entities.GetWeightedRandomElementCode(rnd)
		case entities.BRAND_PASSIVE_DISTORTION:
			item.AsWeapon.Damage += rnd.RandInRange(-1, 1)
			if item.AsWeapon.Damage == 0 {
				item.AsWeapon.Damage = 1
			}
		}
	}
}
