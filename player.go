package main

import "hydras2/entities"

type player struct {
	dungX, dungY int

	currentWeapon     *entities.Item
	currentConsumable *entities.Item
	inventory         []*entities.Item
	keys              map[int]bool

	hitpoints int

	// stats
	strength int // how many items can be carried
	vitality int // how many HP does the player have

	// battlefield-only vars
	x, y int
}

func (p *player) init() {
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
	p.inventory = append(p.inventory, &entities.Item{
		AsConsumable: nil,
		AsWeapon: &entities.ItemWeapon{
			WeaponType: entities.WTYPE_SUBSTRACTOR,
			Damage:     1,
		},
	})
	p.inventory = append(p.inventory, &entities.Item{
		AsConsumable: &entities.ItemConsumable{
			Code:   entities.ITEM_HEAL,
			Amount: 2,
		},
	})
	p.cycleToNextWeapon()
	p.cycleToNextConsumable()
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
	panic("NO ITEM TYPE")
}

func (p *player) cycleToNextWeapon() {
	// shitty code ahead
	selectNextWeapon := p.currentWeapon == nil
	for i := 0; ; i = (i + 1) % len(p.inventory) {
		if p.inventory[i].IsWeapon() && selectNextWeapon {
			p.currentWeapon = p.inventory[i]
			return
		}
		if p.inventory[i] == p.currentWeapon {
			selectNextWeapon = true
		}
	}
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
