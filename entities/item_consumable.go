package entities

import (
	"fmt"
	"github.com/sidav/sidavgorandom/fibrandom"
	"hydras2/text_colors"
)

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
	consumableType        uint8
	name, info, fluffText string
	frequency             int
	defaultEnchantAmount  int
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

func (ic *ItemConsumable) getDescription() string {
	return fmt.Sprintf(
		"%s \n %s \n %s",
		consumablesData[ic.Code].name,
		consumablesData[ic.Code].info,
		text_colors.MakeStringColorTagged(consumablesData[ic.Code].fluffText, []string{"YELLOW"}),
	)
}

var consumablesData = []*ConsumableItemInfo{
	{
		consumableType:       CONSUMABLE_HEAL,
		name:                 "Healing flask",
		frequency:            4,
		defaultEnchantAmount: 2,
		info:                 "Can be used to recover HP.",
		fluffText:            "There is a saying in Hydra Hunters Guild: \"Never throw out your flask, even if emptied\"." +
			" The students of the Guild always think of this as of a stupid moralizing, until they get to know about " +
			"bezoars.",
	},
	{
		consumableType: CONSUMABLE_RETREAT,
		name:           "Escape vial",
		frequency:      2,
		info:           "Use it to flee from combat.",
		fluffText:      "Odd substance, making you just wake up near nearby bonfire. Nobody knows how it works.",
	},
	{
		consumableType: CONSUMABLE_DESTROY_HYDRA,
		name:           "Scroll of destroy hydra",
		frequency:      1,
		info:           "Can be used to destroy hydra.",
		fluffText: "The secret of the scroll was lost long ago. They say that the inventor of this scroll grew overconfident and was later found dead with wounds" +
			"by hydra. It can be that the scroll is not as powerful as it says.",
	},
	{
		consumableType: CONSUMABLE_SHUFFLE_ELEMENT,
		name:           "Scroll of shuffle elements",
		frequency:      1,
		info:           "Can be used to change hydras' elements.",
		fluffText: "Entangles elements' surges of surrounding living beings. Some time ago those scrolls were extremely" +
			" dangerous to use, but time has changed.",
	},
	{
		consumableType: CONSUMABLE_MASS_BISECTION,
		name:           "Scroll of mass bisection",
		info:           "Divides all heads of present hydras by 2.",
		fluffText:      "Sends a wave of awesome bisecting power. It is so strong that it cuts even through the " +
			"power of odd numbers.",
		frequency:      1,
	},
	{
		consumableType: CONSUMABLE_PARALYZE_HYDRAS,
		name:           "Scroll of stagger hydras",
		info:           "Paralyzes all hydras.",
		fluffText:      "This scroll is another of the Hydra Hunters Guild's weapons experiment result. " +
			"This scroll changes some minor amount of hydra's brain matter into curare. Sadly, it appears that " +
			"this really is not enough to kill them.",
		frequency:      1,
	},
	{
		consumableType: CONSUMABLE_UNELEMENT_HYDRAS,
		name:           "Scroll of purge hydras",
		info:           "Purges all hydras.",
		fluffText:      "Destroys elements' links of surrounding living beings. " +
			"Some time ago those scrolls were extremely dangerous to use, but time has changed.",
		frequency:      1,
	},
	{
		consumableType: CONSUMABLE_MERGE_HYDRAS,
		name:           "Scroll of merge hydras",
		info:           "Merges all hydras.",
		fluffText:      "This scroll is another of the Hydra Hunters Guild's weapons experiment result. " +
			"Grinds all hydras' flesh into a mess of living cells cocktail. Sadly, it appears that this mess will" +
			" always reconstruct itself into another hydra in matter of seconds.",
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
