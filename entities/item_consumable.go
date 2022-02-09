package entities

import "github.com/sidav/sidavgorandom/fibrandom"

const (
	ITEM_HEAL = iota
	ITEM_RETREAT
	TOTAL_ITEM_TYPES_NUMBER // for generators
)

type ItemConsumable struct {
	Code int
	Amount int
}

type ConsumableItemInfo struct {
	consumableType uint8
	name, info     string
	frequency      int
}

func GetWeightedRandomConsumableCode(rnd *fibrandom.FibRandom) int {
	return rnd.SelectRandomIndexFromWeighted (
		TOTAL_ITEM_TYPES_NUMBER,
		func(x int) int {return consumablesData[x].frequency},
	)
}

var consumablesData = []*ConsumableItemInfo{
	{
		consumableType: ITEM_HEAL,
		name:           "Healing flask",
		frequency:      2,
		info:           "Can be used to recover HP.",
	},
	{
		consumableType: ITEM_RETREAT,
		name:           "Escape vial",
		frequency:      1,
		info:           "Use it to flee from combat.",
	},
	//{
	//	consumableType: ITEM_ENCHANTER,
	//	name:           "Glyph of enchant weapon",
	//	frequency:      3,
	//	info:           "Can be used to increase weapon damage.",
	//},
	//{
	//	consumableType: ITEM_DESTROY_HYDRA,
	//	name:           "Glyph of destroy hydra",
	//	frequency:      1,
	//	info:           "Can be used to destroy hydra.",
	//},
	//{
	//	consumableType: ITEM_CONFUSE_HYDRA,
	//	name:           "Glyph of confuse hydra",
	//	frequency:      1,
	//	info:           "Can be used to confuse hydra.",
	//},
	//{
	//	consumableType: ITEM_MASS_CONFUSION,
	//	name:           "Scroll of confused party",
	//	info:           "Can be used to confuse all enemies.",
	//	frequency:      1,
	//},
	//{
	//	consumableType: ITEM_INCREASE_HP,
	//	name:           "Potion of vitality",
	//	info:           "Permanently increases your maximum HP and heals you.",
	//	frequency:      3,
	//},
	//{
	//	consumableType: ITEM_STRENGTH,
	//	name:           "Potion of strength",
	//	info:           "Permanently increases your inventory size.",
	//	frequency:      3,
	//},
	//{
	//	consumableType: ITEM_CHANGE_ELEMENT_RANDOM,
	//	name:           "Glyph of randomize element",
	//	info:           "Changes hydra's or Item's element.",
	//	frequency:      1,
	//},
	//{
	//	consumableType: ITEM_CHANGE_ELEMENT_SPECIFIC,
	//	name:           "Glyph of ", // not an error!
	//	info:           "Changes hydra's or Item's element.",
	//	frequency:      1,
	//},
	//{
	//	consumableType: ITEM_BRANDING_RANDOM,
	//	name:           "Glyph of imbue random brand",
	//	info:           "Used to imbue a random brand onto Item.",
	//	frequency:      1,
	//},
	//{
	//	consumableType: ITEM_BRANDING_SPECIFIC,
	//	name:           "Glyph of imbue ", // not an error
	//	info:           "Used to imbue a specific brand onto Item.",
	//	frequency:      1,
	//},
	//{
	//	consumableType: ITEM_UNELEMENT_ENEMIES,
	//	name:           "Scroll of nullification",
	//	info:           "Permanently makes all present enemies non-elemental.",
	//	frequency:      1,
	//},
	//{
	//	consumableType: ITEM_DECAPITATION,
	//	name:           "Scroll of mass bisection",
	//	info:           "Divides all heads of presen hydras by 2.",
	//	frequency:      1,
	//},
	//{
	//	consumableType: ITEM_IMPROVE_BRAND,
	//	name:           "Glyph of improve brand",
	//	info:           "Can be used to improve branded (magic) items",
	//	frequency:      1,
	//},
	//{
	//	consumableType: ITEM_MERGE_HYDRAS_INTO_ONE,
	//	name:           "Scroll of merge hydras",
	//	info:           "Used to make a single hydra from many.",
	//	frequency:      1,
	//},
	//{
	//	consumableType: ITEM_AMMO,
	//	name:           "Charge crystals",
	//	info:           "Used with chargeable items.",
	//	frequency:      2,
	//},
}
