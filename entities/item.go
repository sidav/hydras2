package entities

import "github.com/sidav/sidavgorandom/fibrandom"

type Item struct {
	AsConsumable *ItemConsumable
	AsWeapon     *ItemWeapon
}

func (i *Item) IsConsumable() bool {
	return i.AsConsumable != nil
}

func (i *Item) IsWeapon() bool {
	return i.AsWeapon != nil
}

func (i *Item) GetName() string {
	if i.AsConsumable != nil {
		return consumablesData[i.AsConsumable.Code].name
	}
	if i.AsWeapon != nil {
		return i.AsWeapon.GetName()
	}
	return "NO NAME"
}

func (i *Item) IsStackable() bool {
	return i.IsConsumable()
}

func GenerateRandomItem(rnd *fibrandom.FibRandom) *Item {
	if rnd.OneChanceFrom(3) {
		return &Item{
			AsWeapon: GenerateRandomItemWeapon(rnd),
		}
	}
	return &Item{
		AsConsumable: &ItemConsumable{
			Code:   GetWeightedRandomConsumableCode(rnd),
			Amount: 1,
		},
		AsWeapon:     nil,
	}
}
