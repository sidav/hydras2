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
	BRAND_PASSIVE_DISTORTION
	BRAND_DOUBLE_STRIKE
	BRAND_SAFE_DOUBLE_STRIKE
	BRAND_FAST_STRIKE
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
		canBeOnWeapon: true,
		canBeOnRing:   false,
		isActivatable: false,
		name:          "instability",
		info:          "Changes its element each turn randomly.",
		frequency:     1,
	},
	BRAND_PASSIVE_DISTORTION: {
		canBeOnWeapon: true,
		name:          "distortion",
		info:          "Changes its damage each turn randomly.",
		frequency:     1,
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
		defaultEnchantAmount: 1,
		maxEnchantAmount:     5,
		frequency:            1,
	},
}

func GetWeightedRandomBrandCode(rnd *fibrandom.FibRandom) int {
	return rnd.SelectRandomIndexFromWeighted(len(BrandsTable), func(x int) int { return BrandsTable[x].frequency })
}
