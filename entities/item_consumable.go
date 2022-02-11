package entities

import "github.com/sidav/sidavgorandom/fibrandom"

const (
	CONSUMABLE_HEAL = iota
	CONSUMABLE_RETREAT
	CONSUMABLE_DESTROY_HYDRA
	CONSUMABLE_SHUFFLE_ELEMENT
	CONSUMABLE_MASS_BISECTION
	CONSUMABLE_PARALYZE_HYDRAS
	CONSUMABLE_UNELEMENT_HYDRAS
	CONSUMABLE_MERGE_HYDRAS
)

type ItemConsumable struct {
	Code          int
	Amount        int
	EnchantAmount int
}

type ConsumableItemInfo struct {
	consumableType       uint8
	name, info           string
	frequency            int
	defaultEnchantAmount int
}

func GetWeightedRandomConsumableCode(rnd *fibrandom.FibRandom) int {
	return rnd.SelectRandomIndexFromWeighted(
		len(consumablesData),
		func(x int) int { return consumablesData[x].frequency },
	)
}

func (ic *ItemConsumable) GetDefaultEnchantAmount() int {
	return consumablesData[ic.Code].defaultEnchantAmount
}

var consumablesData = []*ConsumableItemInfo{
	{
		consumableType:       CONSUMABLE_HEAL,
		name:                 "Healing flask",
		frequency:            4,
		defaultEnchantAmount: 2,
		info:                 "Can be used to recover HP.",
	},
	{
		consumableType: CONSUMABLE_RETREAT,
		name:           "Escape vial",
		frequency:      2,
		info:           "Use it to flee from combat.",
	},
	{
		consumableType: CONSUMABLE_DESTROY_HYDRA,
		name:           "Scroll of destroy hydra",
		frequency:      1,
		info:           "Can be used to destroy hydra.",
	},
	{
		consumableType: CONSUMABLE_SHUFFLE_ELEMENT,
		name:           "Scroll of shuffle elements",
		frequency:      1,
		info:           "Can be used to change hydras' elements.",
	},
	{
		consumableType: CONSUMABLE_MASS_BISECTION,
		name:           "Scroll of mass bisection",
		info:           "Divides all heads of present hydras by 2.",
		frequency:      1,
	},
	{
		consumableType: CONSUMABLE_PARALYZE_HYDRAS,
		name:           "Scroll of stagger hydras",
		info:           "Paralyzes all hydras.",
		frequency:      1,
	},
	{
		consumableType: CONSUMABLE_UNELEMENT_HYDRAS,
		name:           "Scroll of purge hydras",
		info:           "Purges all hydras.",
		frequency:      1,
	},
	{
		consumableType: CONSUMABLE_MERGE_HYDRAS,
		name:           "Scroll of merge hydras",
		info:           "Merges all hydras.",
		frequency:      1,
	},
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
	//	consumableType: ITEM_CHANGE_ELEMENT_SPECIFIC,
	//	name:           "Glyph of ", // not an error!
	//	info:           "Changes hydra's or Item's element.",
	//	frequency:      1,
	//},
	//{
	//	consumableType: ITEM_AMMO,
	//	name:           "Charge crystals",
	//	info:           "Used with chargeable items.",
	//	frequency:      2,
	//},
}
