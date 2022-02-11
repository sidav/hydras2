package main

import (
	"hydras2/entities"
)

type player struct {
	dungX, dungY int

	primaryWeapon     *entities.Item
	secondaryWeapon   *entities.Item
	currentConsumable *entities.Item
	inventory         []*entities.Item
	keys              map[int]bool

	hitpoints int

	souls int

	// stats
	level    int
	strength int // how many items can be carried
	vitality int // how many HP does the player have

	// battlefield-only vars
	x, y          int
	nextTickToAct int
}

func (p *player) init() {
	p.level = 1
	p.strength = 5
	p.vitality = 20
	p.hitpoints = p.getMaxHp()

	p.keys = make(map[int]bool, 0)

	p.inventory = append(p.inventory, &entities.Item{
		AsConsumable: nil,
		AsWeapon: &entities.ItemWeapon{
			WeaponTypeCode: entities.WTYPE_SUBSTRACTOR,
			Damage:         2,
		},
	})
	p.primaryWeapon = p.inventory[0]
	p.inventory = append(p.inventory, &entities.Item{
		AsConsumable: nil,
		AsWeapon: &entities.ItemWeapon{
			WeaponTypeCode: entities.WTYPE_SUBSTRACTOR,
			Damage:         1,
		},
	})
	p.secondaryWeapon = p.inventory[1]
	p.inventory = append(p.inventory, &entities.Item{
		AsConsumable: &entities.ItemConsumable{
			Code:          entities.CONSUMABLE_HEAL,
			EnchantAmount: 3,
			Amount:        2,
		},
	})
	p.currentConsumable = p.inventory[2]
}

func (p *player) getMaxHp() int {
	return p.vitality / 2
}

func (p *player) acquireItem(i *entities.Item) {
	if i.IsWeapon() {
		p.inventory = append(p.inventory, i)
		return
	}
	if i.IsConsumable() {
		for ind := range p.inventory {
			if p.inventory[ind].IsConsumable() && p.inventory[ind].AsConsumable.Code == i.AsConsumable.Code {
				p.inventory[ind].AsConsumable.Amount += i.AsConsumable.Amount
				return
			}
		}
		p.inventory = append(p.inventory, i)
		return
	}
	if i.IsMaterial() {
		p.inventory = append(p.inventory, i)
		return
	}
	panic("NO ITEM TYPE")
}

func (p *player) cycleToNextPrimaryWeapon() {
	// shitty code ahead
	selectNextWeapon := p.primaryWeapon == nil
	for i := 0; ; i = (i + 1) % len(p.inventory) {
		if p.inventory[i] == p.secondaryWeapon {
			continue
		}
		if p.inventory[i].IsWeapon() && selectNextWeapon {
			p.primaryWeapon = p.inventory[i]
			return
		}
		if p.inventory[i] == p.primaryWeapon {
			selectNextWeapon = true
		}
	}
}

func (p *player) cycleToNextSecondaryWeapon() {
	// shitty code ahead
	selectNextWeapon := p.secondaryWeapon == nil
	for i := 0; ; i = (i + 1) % len(p.inventory) {
		if p.inventory[i] == p.primaryWeapon {
			continue
		}
		if p.inventory[i].IsWeapon() && selectNextWeapon {
			p.secondaryWeapon = p.inventory[i]
			return
		}
		if p.inventory[i] == p.secondaryWeapon {
			selectNextWeapon = true
		}
	}
}

func (p *player) swapWeapons() {
	t := p.primaryWeapon
	p.primaryWeapon = p.secondaryWeapon
	p.secondaryWeapon = t
}

func (p *player) cycleToNextConsumable() {
	// shitty code ahead
	selectNextItem := p.currentConsumable == nil
	for i := 0; ; i = (i + 1) % len(p.inventory) {
		if p.inventory[i].IsConsumable() && selectNextItem {
			p.currentConsumable = p.inventory[i]
			return
		}
		if p.inventory[i] == p.currentConsumable {
			selectNextItem = true
		}
	}
}

func (p *player) removeItemFromInventory(itmToRemove *entities.Item) {
	for i, item := range p.inventory {
		if item == itmToRemove {
			if item.IsConsumable() {
				item.AsConsumable.Amount--
				if item.AsConsumable.Amount > 0 || item.AsConsumable.EnchantAmount > item.AsConsumable.GetDefaultEnchantAmount() {
					return
				}
			}
			p.inventory = append(p.inventory[:i], p.inventory[i+1:]...)
			if p.currentConsumable == itmToRemove {
				p.currentConsumable = nil
				p.cycleToNextConsumable()
			}
			p.sortInventory()
			return
		}
	}
	panic("No such item to remove!")
}

func (p *player) getAllItemsOfType(itype int) []*entities.Item {
	var items []*entities.Item
	for _, i := range p.inventory {
		t, _ := i.GetTypeAndCode()
		if t == itype {
			items = append(items, i)
		}
	}
	return items
}

func (p *player) getAllItemsOfTypeAndCode(itype, icode int) []*entities.Item {
	var items []*entities.Item
	for _, i := range p.inventory {
		t, c := i.GetTypeAndCode()
		if t == itype && c == icode {
			items = append(items, i)
		}
	}
	return items
}

func (p *player) getAllWeapons() []*entities.Item {
	return p.getAllItemsOfType(entities.ITEM_WEAPON)
}

func (p *player) getAllMaterials() []*entities.Item {
	return p.getAllItemsOfType(entities.ITEM_MATERIAL)
}

func (p *player) sortInventory() {
	entities.SortItemsArray(p.inventory)
}
