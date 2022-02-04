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
	return "NO NAME"
}
