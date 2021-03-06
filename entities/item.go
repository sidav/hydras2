package entities

import (
	"fmt"
	"github.com/sidav/sidavgorandom/fibrandom"
	"sort"
)

const (
	ITEM_WEAPON = iota
	ITEM_CONSUMABLE
	ITEM_MATERIAL
)

type Item struct {
	AsConsumable *ItemConsumable
	AsWeapon     *ItemWeapon
	AsMaterial   *ItemMaterial
}

func SortItemsArray(iArr []*Item) {
	sort.Slice(iArr,
		func(i, j int) bool {
			if iArr[i].IsWeapon() && !iArr[j].IsWeapon() {
				return true
			}
			if iArr[i].IsWeapon() && iArr[j].IsWeapon() {
				if iArr[i].AsWeapon.WeaponTypeCode < iArr[j].AsWeapon.WeaponTypeCode {
					return true
				}
				if iArr[i].AsWeapon.Damage > iArr[j].AsWeapon.Damage {
					return true
				}
			}
			if iArr[i].IsConsumable() && !iArr[j].IsWeapon() {
				return true
			}
			if iArr[i].IsMaterial() && iArr[j].IsMaterial() {
				return iArr[i].AsMaterial.Code < iArr[j].AsMaterial.Code
			}
			return false
		},
	)
}

func (i *Item) IsConsumable() bool {
	return i.AsConsumable != nil
}

func (i *Item) IsWeapon() bool {
	return i.AsWeapon != nil
}

func (i *Item) IsMaterial() bool {
	return i.AsMaterial != nil
}

func (i *Item) GetTypeAndCode() (int, int) {
	if i.IsWeapon() {
		return ITEM_WEAPON, i.AsWeapon.WeaponTypeCode
	}
	if i.IsConsumable() {
		return ITEM_CONSUMABLE, i.AsConsumable.Code
	}
	if i.IsMaterial() {
		return ITEM_MATERIAL, i.AsMaterial.Code
	}
	panic("What the heck is this item?!")
}

func (i *Item) GetName() string {
	if i.AsConsumable != nil {
		name := consumablesData[i.AsConsumable.Code].name
		if i.AsConsumable.EnchantAmount > consumablesData[i.AsConsumable.Code].defaultEnchantAmount {
			name += fmt.Sprintf(" +%d", i.AsConsumable.EnchantAmount)
		}
		return name
	}
	if i.AsWeapon != nil {
		return i.AsWeapon.GetName()
	}
	if i.AsMaterial != nil {
		return i.AsMaterial.GetName()
	}
	panic("No item name!")
}

func (i *Item) GetDescription() string {
	if i.IsConsumable() {
		return i.AsConsumable.getDescription()
	}
	if i.IsWeapon() {
		return i.AsWeapon.getDescription()
	}
	if i.IsMaterial() {
		return i.AsMaterial.getDescription()
	}
	panic("NO DESCRIPTION")
}

func (i *Item) IsStackable() bool {
	return i.IsConsumable()
}

func GenerateRandomItem(rnd *fibrandom.FibRandom) *Item {
	typeFrequencies := []int{2, 3, 3}
	whatToGen := rnd.SelectRandomIndexFromWeighted(len(typeFrequencies), func(x int) int { return typeFrequencies[x] })
	switch whatToGen {
	case 0: // weapon
		return &Item{
			AsWeapon: GenerateRandomItemWeapon(rnd),
		}
	case 1: // consumable
		code := GetWeightedRandomConsumableCode(rnd)
		return &Item{
			AsConsumable: &ItemConsumable{
				Code:          code,
				EnchantAmount: consumablesData[code].defaultEnchantAmount,
				Amount:        1,
			},
		}
	case 2: // material
		return &Item{
			AsMaterial: GenerateRandomMaterial(rnd),
		}
	default:
		panic("Wrong generate")
	}
}
