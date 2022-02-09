package main

import "hydras2/entities"

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
	x, y int
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
			WeaponType: entities.WTYPE_SUBSTRACTOR,
			Damage:     2,
		},
	})
	p.primaryWeapon = p.inventory[0]
	p.inventory = append(p.inventory, &entities.Item{
		AsConsumable: nil,
		AsWeapon: &entities.ItemWeapon{
			WeaponType: entities.WTYPE_SUBSTRACTOR,
			Damage:     1,
		},
	})
	p.secondaryWeapon = p.inventory[1]
	p.inventory = append(p.inventory, &entities.Item{
		AsConsumable: &entities.ItemConsumable{
			Code:   entities.ITEM_HEAL,
			Amount: 2,
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
			p.inventory = append(p.inventory[:i], p.inventory[i+1:]...)
			return
		}
	}
	panic("No such item to remove!")
}

func (p *player) getAllWeapons() []*entities.Item {
	var weps []*entities.Item
	for _, i := range p.inventory {
		if i.IsWeapon() {
			weps = append(weps, i)
		}
	}
	return weps
}

func (p *player) getAllMaterials() []*entities.Item {
	var mats []*entities.Item
	for _, i := range p.inventory {
		if i.IsMaterial() {
			mats = append(mats, i)
		}
	}
	return mats
}

func (p *player) getAllMaterialsOfCode(code int) []*entities.Item {
	var mats []*entities.Item
	for _, i := range p.inventory {
		if i.IsMaterial() && i.AsMaterial.Code == code {
			mats = append(mats, i)
		}
	}
	return mats
}
