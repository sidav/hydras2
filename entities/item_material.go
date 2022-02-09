package entities

import (
	"fmt"
	"github.com/sidav/sidavgorandom/fibrandom"
)

const (
	MATERIAL_CLEAR_BRAND = iota
	MATERIAL_IMBUE_BRAND
	MATERIAL_IMPROVE_BRAND
	MATERIAL_APPLY_ELEMENT
	MATERIAL_ENCHANT
	MATERIAL_ENCHANT_CONSUMABLE
)

type ItemMaterial struct {
	Code           int
	ImbuesBrand    *Brand
	AppliesElement *Element
	EnchantAmount  int
}

func (im *ItemMaterial) GetName() string {
	switch im.Code {
	case MATERIAL_CLEAR_BRAND:
		return "Clear quartz"
	case MATERIAL_IMBUE_BRAND:
		return "Glyph of " + im.ImbuesBrand.GetName()
	case MATERIAL_IMPROVE_BRAND:
		return "Stone of power"
	case MATERIAL_APPLY_ELEMENT:
		if im.AppliesElement.Code == ELEMENT_NONE {
			return "Purging gem"
		} else {
			return im.AppliesElement.GetName() + " gem"
		}
	case MATERIAL_ENCHANT:
		if im.EnchantAmount > 0 {
			return fmt.Sprintf("Sharpener (+%d)", im.EnchantAmount)
		} else {
			return fmt.Sprintf("Sharpener (%d)", im.EnchantAmount)
		}
	case MATERIAL_ENCHANT_CONSUMABLE:
		return "Bezoar"
	}
	panic("No GetName")
}

func GenerateRandomMaterial(rnd *fibrandom.FibRandom) *ItemMaterial {
	typeFrequencies := []int{2, 1, 1, 1, 3, 1}
	whatToGen := rnd.SelectRandomIndexFromWeighted(len(typeFrequencies), func(x int) int { return typeFrequencies[x] })
	switch whatToGen {
	case 0:
		return &ItemMaterial{
			Code: MATERIAL_CLEAR_BRAND,
		}
	case 1:
		return &ItemMaterial{
			Code:        MATERIAL_IMBUE_BRAND,
			ImbuesBrand: &Brand{rnd.Rand(len(BrandsTable))},
		}
	case 2:
		return &ItemMaterial{
			Code: MATERIAL_IMPROVE_BRAND,
		}
	case 3:
		return &ItemMaterial{
			Code:           MATERIAL_APPLY_ELEMENT,
			AppliesElement: &Element{GetWeightedRandomElementCode(rnd)},
		}
	case 4:
		enchantProbabilities := []int{1, 2, 0, 3, 1}
		enchantAmount := rnd.SelectRandomIndexFromWeighted(len(enchantProbabilities), func(x int) int { return enchantProbabilities[x] })
		return &ItemMaterial{
			Code:          MATERIAL_ENCHANT,
			EnchantAmount: enchantAmount-len(enchantProbabilities)/2,
		}
	case 5:
		return &ItemMaterial{
			Code: MATERIAL_ENCHANT_CONSUMABLE,
		}
	default:
		panic("Wrong generate")
	}
}
