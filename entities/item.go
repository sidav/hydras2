package entities

import "github.com/sidav/sidavgorandom/fibrandom"

type Item struct {
	AsConsumable *ItemConsumable
	AsWeapon     *ItemWeapon
	AsMaterial   *ItemMaterial
}

func (i *Item) IsConsumable() bool {
	return i.AsConsumable != nil
}

func (i *Item) IsWeapon() bool {
	return i.AsWeapon != nil
}

func (i *Item) IsMaterial() bool {
	return i.AsMaterial != nil
}

func (i *Item) GetName() string {
	if i.AsConsumable != nil {
		return consumablesData[i.AsConsumable.Code].name
	}
	if i.AsWeapon != nil {
		return i.AsWeapon.GetName()
	}
	if i.AsMaterial != nil {
		return i.AsMaterial.GetName()
	}
	panic("No item name!")
}

func (i *Item) IsStackable() bool {
	return i.IsConsumable()
}

func GenerateRandomItem(rnd *fibrandom.FibRandom) *Item {
	typeFrequencies := []int{1, 2, 1}
	whatToGen := rnd.SelectRandomIndexFromWeighted(len(typeFrequencies), func(x int) int { return typeFrequencies[x] })
	switch whatToGen {
	case 0: // weapon
		return &Item{
			AsWeapon: GenerateRandomItemWeapon(rnd),
		}
	case 1: // consumable
		return &Item{
			AsConsumable: &ItemConsumable{
				Code:   GetWeightedRandomConsumableCode(rnd),
				Amount: 1,
			},
		}
	case 2: // material
		return &Item{
			AsMaterial: GenerateRandomMaterial(rnd),
		}
	default:
		panic("Wrong generate")
	}
}
