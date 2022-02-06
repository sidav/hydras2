package main

type item struct {
	asConsumable *consumableItemInfo
	asWeapon     *itemWeapon
}

func (i *item) isConsumable() bool {
	return i.asConsumable != nil
}

func (i *item) getName() string {
	if i.asConsumable != nil {
		return i.asConsumable.name
	}
	if i.asWeapon != nil {
		return i.asWeapon.getName()
	}
	return "NO NAME"
}

func generateRandomItem() *item {
	if rnd.OneChanceFrom(3) {
		return &item{
			asWeapon:     generateRandomItemWeapon(),
		}
	}
	return &item{
		asConsumable: consumablesData[getWeightedRandomConsumableCode()],
		asWeapon:     nil,
	}
}
