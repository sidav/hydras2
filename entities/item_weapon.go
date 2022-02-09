package entities

import (
	"fmt"
	"github.com/sidav/sidavgorandom/fibrandom"
	"hydras2/text_colors"
)

const (
	WTYPE_SUBSTRACTOR = iota
	WTYPE_DIVISOR
	WTYPE_LOGARITHMER
)

type WeaponTypeStaticData struct {
	WeaponTypeCode         int
	Frequency              int
	MinDamageForGeneration int
}

var weaponsStaticData = []*WeaponTypeStaticData{
	{
		WeaponTypeCode:         WTYPE_SUBSTRACTOR,
		Frequency:              4,
		MinDamageForGeneration: 1,
	},
	{
		WeaponTypeCode:         WTYPE_DIVISOR,
		Frequency:              2,
		MinDamageForGeneration: 2,
	},
	{
		WeaponTypeCode:         WTYPE_LOGARITHMER,
		Frequency:              1,
		MinDamageForGeneration: 2,
	},
}

type ItemWeapon struct {
	WeaponTypeCode int
	WeaponElement  Element
	Brand          *Brand
	Damage         int
}

func GenerateRandomItemWeapon(rnd *fibrandom.FibRandom) *ItemWeapon {
	index := rnd.SelectRandomIndexFromWeighted(len(weaponsStaticData), func(i int) int { return weaponsStaticData[i].Frequency })
	var brand *Brand
	if rnd.OneChanceFrom(3) { // TODO: actual chance
		brand = &Brand{
			Code: rnd.Rand(len(BrandsTable)),
		}
	}
	iw := ItemWeapon{
		WeaponTypeCode: index,
		WeaponElement:  Element{GetWeightedRandomElementCode(rnd)},
		Brand:          brand,
		Damage:         rnd.RandInRange(weaponsStaticData[index].MinDamageForGeneration, weaponsStaticData[index].MinDamageForGeneration+2),
	}
	return &iw
}

func (w *ItemWeapon) GetName() string {
	name := w.WeaponElement.GetName()
	if len(name) > 0 {
		name += " "
	}
	switch w.WeaponTypeCode {
	case WTYPE_SUBSTRACTOR:
		name += fmt.Sprintf("-%d Substractor", w.Damage)
	case WTYPE_DIVISOR:
		switch w.Damage {
		case 2:
			name += "Bisector"
		case 3:
			name += "Trisector"
		case 10:
			name += "Decimator"
		default:
			name += fmt.Sprintf("/%d Divisor", w.Damage)
		}
	case WTYPE_LOGARITHMER:
		name += fmt.Sprintf("%d-logarithmer", w.Damage)
	default:
		panic("No ItemWeapon name")
	}
	if w.Brand != nil {
		name += " of " + w.Brand.GetName()
	}
	return text_colors.MakeStringColorTagged(name, w.WeaponElement.GetColorTags())
}

func (iw *ItemWeapon) GetDamageOnHeads(heads int) int {
	switch iw.WeaponTypeCode {
	case WTYPE_SUBSTRACTOR:
		if heads < iw.Damage {
			return 0
		}
		return iw.Damage
	case WTYPE_DIVISOR:
		if heads%iw.Damage > 0 {
			return 0
		}
		return heads - (heads / iw.Damage)
	case WTYPE_LOGARITHMER:
		panic("No logarithmer implemented!")
	}
	return 0
}
