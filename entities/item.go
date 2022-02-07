package entities

import "github.com/sidav/sidavgorandom/fibrandom"

type Item struct {
	AsConsumable *consumableItemInfo
	AsWeapon     *ItemWeapon
}

func (i *Item) IsConsumable() bool {
	return i.AsConsumable != nil
}

func (i *Item) GetName() string {
	if i.AsConsumable != nil {
		return i.AsConsumable.name
	}
	if i.AsWeapon != nil {
		return i.AsWeapon.GetName()
	}
	return "NO NAME"
}

func GenerateRandomItem(rnd *fibrandom.FibRandom) *Item {
	if rnd.OneChanceFrom(3) {
		return &Item{
			AsWeapon: GenerateRandomItemWeapon(rnd),
		}
	}
	return &Item{
		AsConsumable: consumablesData[GetWeightedRandomConsumableCode(rnd)],
		AsWeapon:     nil,
	}
}
