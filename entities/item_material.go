package entities

import (
	"fmt"
	"github.com/sidav/sidavgorandom/fibrandom"
)

const (
	IMTYPE_CLEAR_BRAND = iota
	IMTYPE_IMBUE_BRAND
	IMTYPE_IMPROVE_BRAND
	IMTYPE_APPLY_ELEMENT
	IMTYPE_ENCHANT
)

type ItemMaterial struct {
	Code           int
	ImbuesBrand    *Brand
	AppliesElement *Element
	EnchantAmount  int
}

func (im *ItemMaterial) GetName() string {
	switch im.Code {
	case IMTYPE_CLEAR_BRAND:
		return "Clear quartz"
	case IMTYPE_IMBUE_BRAND:
		return "Glyph of " + im.ImbuesBrand.GetName()
	case IMTYPE_IMPROVE_BRAND:
		return "Stone of power"
	case IMTYPE_APPLY_ELEMENT:
		return im.AppliesElement.GetName() + " gem"
	case IMTYPE_ENCHANT:
		if im.EnchantAmount > 0 {
			return fmt.Sprintf("Sharpener (+%d)", im.EnchantAmount)
		} else {
			return fmt.Sprintf("Sharpener (%d)", im.EnchantAmount)
		}
	}
	panic("No GetName")
}

func GenerateRandomMaterial(rnd *fibrandom.FibRandom) *ItemMaterial {
	typeFrequencies := []int{2, 1, 1, 1, 3}
	whatToGen := rnd.SelectRandomIndexFromWeighted(len(typeFrequencies), func(x int) int { return typeFrequencies[x] })
	switch whatToGen {
	case 0:
		return &ItemMaterial{
			Code:           IMTYPE_CLEAR_BRAND,
		}
	case 1:
		return &ItemMaterial{
			Code: IMTYPE_IMBUE_BRAND,
			ImbuesBrand:    &Brand{rnd.Rand(len(BrandsTable))},
		}
	case 2:
		return &ItemMaterial{
			Code: IMTYPE_IMPROVE_BRAND,
		}
	case 3:
		return &ItemMaterial{
			Code: IMTYPE_APPLY_ELEMENT,
			AppliesElement: &Element{GetWeightedRandomElementCode(rnd)},
		}
	case 4:
		enchantProbabilities := []int{1, 2, 0, 3, 1}
		enchantAmount := rnd.SelectRandomIndexFromWeighted(len(enchantProbabilities), func(x int) int { return enchantProbabilities[x] })
		return &ItemMaterial{
			Code: IMTYPE_ENCHANT,
			EnchantAmount: enchantAmount-len(enchantProbabilities)/2,
		}
	default:
		panic("Wrong generate")
	}
}
