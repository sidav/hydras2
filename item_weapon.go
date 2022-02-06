package main

import "fmt"

type weaponType uint8

const (
	WTYPE_SUBSTRACTOR weaponType = iota
	WTYPE_DIVISOR
	WTYPE_LOGARITHMER
)

type weaponTypeStaticData struct {
	wtype                  weaponType
	frequency              int
	minDamageForGeneration int
}

var weaponsStaticData = []*weaponTypeStaticData{
	{
		wtype:                  WTYPE_SUBSTRACTOR,
		frequency:              4,
		minDamageForGeneration: 1,
	},
	{
		wtype:                  WTYPE_DIVISOR,
		frequency:              2,
		minDamageForGeneration: 2,
	},
	{
		wtype:                  WTYPE_LOGARITHMER,
		frequency:              1,
		minDamageForGeneration: 2,
	},
}

type itemWeapon struct {
	weaponType weaponType
	damage     int
}

func generateRandomItemWeapon() *itemWeapon {
	index := rnd.SelectRandomIndexFromWeighted(len(weaponsStaticData), func(i int) int { return weaponsStaticData[i].frequency })
	iw := itemWeapon{
		weaponType: weaponType(index),
		damage:     rnd.RandInRange(weaponsStaticData[index].minDamageForGeneration, weaponsStaticData[index].minDamageForGeneration + 2),
	}
	return &iw
}

func (w *itemWeapon) getName() string {
	name := ""
	switch w.weaponType {
	case WTYPE_SUBSTRACTOR:
		name += fmt.Sprintf("-%d Substractor", w.damage)
	case WTYPE_DIVISOR:
		switch w.damage {
		case 2:
			name += "Bisector"
		case 3:
			name += "Trisector"
		case 10:
			name += "Decimator"
		default:
			name += fmt.Sprintf("/%d Divisor", w.damage)
		}
	case WTYPE_LOGARITHMER:
		name += fmt.Sprintf("%d-logarithmer", w.damage)
	default:
		panic("No itemWeapon name")
	}
	return name
}

