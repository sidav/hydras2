package entities

import (
	"fmt"
	"github.com/sidav/sidavgorandom/fibrandom"
)

type weaponType uint8

const (
	WTYPE_SUBSTRACTOR weaponType = iota
	WTYPE_DIVISOR
	WTYPE_LOGARITHMER
)

type WeaponTypeStaticData struct {
	Wtype                  weaponType
	Frequency              int
	MinDamageForGeneration int
}

var weaponsStaticData = []*WeaponTypeStaticData{
	{
		Wtype:                  WTYPE_SUBSTRACTOR,
		Frequency:              4,
		MinDamageForGeneration: 1,
	},
	{
		Wtype:                  WTYPE_DIVISOR,
		Frequency:              2,
		MinDamageForGeneration: 2,
	},
	{
		Wtype:                  WTYPE_LOGARITHMER,
		Frequency:              1,
		MinDamageForGeneration: 2,
	},
}

type ItemWeapon struct {
	WeaponType weaponType
	Damage     int
}

func GenerateRandomItemWeapon(rnd *fibrandom.FibRandom) *ItemWeapon {
	index := rnd.SelectRandomIndexFromWeighted(len(weaponsStaticData), func(i int) int { return weaponsStaticData[i].Frequency })
	iw := ItemWeapon{
		WeaponType: weaponType(index),
		Damage:     rnd.RandInRange(weaponsStaticData[index].MinDamageForGeneration, weaponsStaticData[index].MinDamageForGeneration+ 2),
	}
	return &iw
}

func (w *ItemWeapon) GetName() string {
	name := ""
	switch w.WeaponType {
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
	return name
}

func (iw *ItemWeapon) GetDamageOnHeads(heads int) int {
	switch iw.WeaponType {
	case WTYPE_SUBSTRACTOR:
		if heads < iw.Damage {
			return 0
		}
		return iw.Damage
	case WTYPE_DIVISOR:
		if iw.Damage% heads > 0 {
			return 0
		}
		return heads-(heads/iw.Damage)
	case WTYPE_LOGARITHMER:
		panic("No logarithmer implemented!")
	}
	return 0
}
