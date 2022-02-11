package entities

import (
	"fmt"
	"github.com/sidav/sidavgorandom/fibrandom"
)

type Brand struct {
	Code          int
	EnchantAmount int
}

func GenerateRandomBrand(rnd *fibrandom.FibRandom) *Brand {
	code := GetWeightedRandomBrandCode(rnd)
	return &Brand{
		Code:          code,
		EnchantAmount: BrandsTable[code].defaultEnchantAmount,
	}
}

func (b *Brand) GetName() string {
	if BrandsTable[b.Code].isEnchantable {
		return fmt.Sprintf("%s +%d", BrandsTable[b.Code].name, b.EnchantAmount)
	}
	return BrandsTable[b.Code].name
}

func (b *Brand) Improve() { // enchant or get improved
	if BrandsTable[b.Code].isEnchantable && b.EnchantAmount < BrandsTable[b.Code].maxEnchantAmount {
		b.EnchantAmount++
	}
	if BrandsTable[b.Code].upgradedVersionCode != 0 {
		b.Code = BrandsTable[b.Code].upgradedVersionCode
	}
}

const (
	BRAND_PASSIVE_ELEMENTS_SHIFTING = iota
	BRAND_PASSIVE_RARE_ELEMENTS_SHIFTING
	BRAND_PASSIVE_DISTORTION
	BRAND_PASSIVE_SMARTER_DISTORTION
	BRAND_DOUBLE_STRIKE
	BRAND_SAFE_DOUBLE_STRIKE
	BRAND_FAST_STRIKE
	BRAND_VAMPIRISM
	BRAND_BETTER_VAMPIRISM
)

type BrandData struct {
	canBeOnWeapon        bool
	canBeOnRing          bool
	isActivatable        bool
	isEnchantable        bool
	defaultEnchantAmount int
	maxEnchantAmount     int
	upgradedVersionCode  int
	name, info           string
	frequency            int
}

var BrandsTable = map[int]*BrandData{
	BRAND_PASSIVE_ELEMENTS_SHIFTING: {
		canBeOnWeapon:       true,
		canBeOnRing:         false,
		isActivatable:       false,
		name:                "elements",
		info:                "Changes its element each turn randomly.",
		upgradedVersionCode: BRAND_PASSIVE_RARE_ELEMENTS_SHIFTING,
		frequency:           1,
	},
	BRAND_PASSIVE_RARE_ELEMENTS_SHIFTING: {
		canBeOnWeapon: true,
		canBeOnRing:   false,
		isActivatable: false,
		name:          "space",
		info:          "Changes its element each turn randomly.",
		frequency:     0,
	},
	BRAND_PASSIVE_DISTORTION: {
		canBeOnWeapon:       true,
		name:                "distortion",
		info:                "Changes its damage each turn randomly.",
		upgradedVersionCode: BRAND_PASSIVE_SMARTER_DISTORTION,
		frequency:           1,
	},
	BRAND_PASSIVE_SMARTER_DISTORTION: {
		canBeOnWeapon: true,
		name:          "attuned distortion",
		info:          "Changes its damage each turn randomly.",
		frequency:     0,
	},
	BRAND_DOUBLE_STRIKE: {
		canBeOnWeapon:       true,
		name:                "double strike",
		info:                "Hits twice.",
		upgradedVersionCode: BRAND_SAFE_DOUBLE_STRIKE,
		frequency:           1,
	},
	BRAND_SAFE_DOUBLE_STRIKE: {
		canBeOnWeapon: true,
		name:          "swift strike",
		info:          "Hits twice if needed.",
		frequency:     0,
	},
	BRAND_FAST_STRIKE: {
		canBeOnWeapon:        true,
		name:                 "fast strike",
		info:                 "Hits faster.",
		isEnchantable:        true,
		defaultEnchantAmount: 3,
		maxEnchantAmount:     5,
		frequency:            1,
	},
	BRAND_VAMPIRISM: {
		canBeOnWeapon:        true,
		isActivatable:        false,
		isEnchantable:        false,
		upgradedVersionCode:  BRAND_BETTER_VAMPIRISM,
		name:                 "vampirism",
		info:                 "10% chance to gain 1 health on hit",
		frequency:            1,
	},
	BRAND_BETTER_VAMPIRISM: {
		canBeOnWeapon:        true,
		isActivatable:        false,
		isEnchantable:        false,
		upgradedVersionCode:  0,
		name:                 "dark craving",
		info:                 "10% chance to gain 2 health on hit",
		frequency:            0,
	},
}

func GetWeightedRandomBrandCode(rnd *fibrandom.FibRandom) int {
	return rnd.SelectRandomIndexFromWeighted(len(BrandsTable), func(x int) int { return BrandsTable[x].frequency })
}
