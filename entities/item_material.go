package entities

import (
	"fmt"
	"github.com/sidav/sidavgorandom/fibrandom"
	"hydras2/text_colors"
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
		return "Cleancing glyph"
	case MATERIAL_IMBUE_BRAND:
		return "Glyph of " + im.ImbuesBrand.GetName()
	case MATERIAL_IMPROVE_BRAND:
		return "Stoneglyph of power"
	case MATERIAL_APPLY_ELEMENT:
		if im.AppliesElement.Code == ELEMENT_NONE {
			return "Purging gem"
		} else {
			return text_colors.MakeStringColorTagged(im.AppliesElement.GetName() + " gem", im.AppliesElement.GetColorTags())
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

func (im *ItemMaterial) getDescription() string {
	switch im.Code {
	case MATERIAL_CLEAR_BRAND:
		return "Cleansing glyph. \n It removes any brand on a weapon. \n" +
			"It is said in Book of Arkanists that any glyph is bound to be purged someday."
	case MATERIAL_IMBUE_BRAND:
		return "Glyph of " + im.ImbuesBrand.GetName() + ". \nThe powerful magic sigil, which will imbue a strong magic " +
			"on any weapon. " + im.ImbuesBrand.getDescription()
	case MATERIAL_IMPROVE_BRAND:
		return "Stoneglyph of power. \nIt is a magic symbol of unknown origin. It is said to be able to enforce any brand with dark power."
	case MATERIAL_APPLY_ELEMENT:
		if im.AppliesElement.Code == ELEMENT_NONE {
			return "\nThose gems are growing from bones of elementless hydras. Those gems somehow can unlink elements " +
				"from any soulless item. Scientists of Tarlidorf are still intrigued of this secret."
		} else {
			return text_colors.MakeStringColorTagged(im.AppliesElement.GetName() + " gem", im.AppliesElement.GetColorTags()) +
				"\nThose gems are just always growing in hydras' lairs. This may be tied with their elemental origin."
		}
	case MATERIAL_ENCHANT:
		str := ""
		if im.EnchantAmount > 0 {
			str = fmt.Sprintf("Sharpener (+%d)", im.EnchantAmount)
		} else {
			str = fmt.Sprintf("Sharpener (%d)", im.EnchantAmount)
		}
		return str + "\nA good sharpener may reinforce one's weapon, as well as ruin it. Use it wisely."
	case MATERIAL_ENCHANT_CONSUMABLE:
		return "Bezoar. \nAn ancient organic stone with hydra's petrified blood in it. It can give a good share of " +
			"hydras' regeneration. The safest way to get it is through one's healing flask."
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
			ImbuesBrand: GenerateRandomBrand(rnd),
		}
	case 2:
		return &ItemMaterial{
			Code: MATERIAL_IMPROVE_BRAND,
		}
	case 3:
		return &ItemMaterial{
			Code:           MATERIAL_APPLY_ELEMENT,
			AppliesElement: &Element{GetRandomElementCode(rnd)},
		}
	case 4:
		enchantProbabilities := []int{1, 2, 0, 3, 1}
		enchantAmount := rnd.SelectRandomIndexFromWeighted(len(enchantProbabilities), func(x int) int { return enchantProbabilities[x] })
		return &ItemMaterial{
			Code:          MATERIAL_ENCHANT,
			EnchantAmount: enchantAmount - len(enchantProbabilities)/2,
		}
	case 5:
		return &ItemMaterial{
			Code: MATERIAL_ENCHANT_CONSUMABLE,
		}
	default:
		panic("Wrong generate")
	}
}
