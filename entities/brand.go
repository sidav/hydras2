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
)

type BrandType struct {
	canBeOnWeapon bool
	canBeOnRing   bool
	isActivatable bool
	name, info    string
}

var BrandsTable = map[int]*BrandType{
	BRAND_PASSIVE_ELEMENTS_SHIFTING: {
		canBeOnWeapon: true,
		canBeOnRing:   false,
		isActivatable: false,
		name:          "instability",
		info:          "Changes its element each turn randomly.",
	},
	BRAND_PASSIVE_DISTORTION: {
		canBeOnWeapon:        true,
		name:                 "distortion",
		info:                 "Changes its damage each turn randomly.",
	},
}

//func GetWeightedRandomBrandCode(rnd *fibrandom.FibRandom) int {
//	rnd.SelectRandomIndexFromWeighted(len(BrandsTable), func(x int) int {return BrandsTable[x].frequency})
//}
