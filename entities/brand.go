package entities

type Brand struct {
	Code int
}

func (b *Brand) GetName() string {
	return BrandsTable[b.Code].name
}

const (
	BRAND_PASSIVE_ELEMENTS_SHIFTING = iota
	BRAND_PASSIVE_DISTORTION
	BRAND_DOUBLE_STRIKE
	BRAND_FAST_STRIKE
)

type BrandData struct {
	canBeOnWeapon       bool
	canBeOnRing         bool
	isActivatable       bool
	upgradedVersionCode int
	name, info          string
}

var BrandsTable = map[int]*BrandData{
	BRAND_PASSIVE_ELEMENTS_SHIFTING: {
		canBeOnWeapon: true,
		canBeOnRing:   false,
		isActivatable: false,
		name:          "instability",
		info:          "Changes its element each turn randomly.",
	},
	BRAND_PASSIVE_DISTORTION: {
		canBeOnWeapon: true,
		name:          "distortion",
		info:          "Changes its damage each turn randomly.",
	},
	BRAND_DOUBLE_STRIKE: {
		canBeOnWeapon:       true,
		name:                "double strike",
		info:                "Hits twice.",
		upgradedVersionCode: BRAND_FAST_STRIKE,
	},
	//BRAND_FAST_STRIKE: {
	//	canBeOnWeapon: true,
	//	name:          "swift strike",
	//	info:          "Hits faster.",
	//},
}

//func GetWeightedRandomBrandCode(rnd *fibrandom.FibRandom) int {
//	rnd.SelectRandomIndexFromWeighted(len(BrandsTable), func(x int) int {return BrandsTable[x].frequency})
//}
