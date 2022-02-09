package main

import (
	"hydras2/entities"
)

func applyMaterialToItem(mat, item *entities.Item) bool {
	applied := true
	switch mat.AsMaterial.Code {
	case entities.MATERIAL_CLEAR_BRAND:
		item.AsWeapon.Brand = nil
	case entities.MATERIAL_IMBUE_BRAND:
		item.AsWeapon.Brand = mat.AsMaterial.ImbuesBrand
	case entities.MATERIAL_IMPROVE_BRAND:

	case entities.MATERIAL_APPLY_ELEMENT:
		item.AsWeapon.WeaponElement.Code = mat.AsMaterial.AppliesElement.Code
	case entities.MATERIAL_ENCHANT:
		item.AsWeapon.Damage += mat.AsMaterial.EnchantAmount
	case entities.MATERIAL_ENCHANT_CONSUMABLE:
		item.AsConsumable.Enchantment++
	}
	return applied
}
