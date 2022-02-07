package main

import "hydras2/entities"

type player struct {
	dungX, dungY  int
	currentWeapon *entities.Item
	currentItem   *entities.Item
	inventory     []*entities.Item
	hitpoints     int

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
	p.cycleToNextWeapon()
}

func (p *player) getMaxHp() int {
	return p.vitality/2
}

func (p *player) cycleToNextWeapon() {
	// shitty code ahead
	selectNextWeapon := p.currentWeapon == nil
	for i := 0; ; i=(i+1)%len(p.inventory) {
		if p.inventory[i].AsWeapon != nil && selectNextWeapon {
			p.currentWeapon = p.inventory[i]
			return
		}
		if p.inventory[i] == p.currentWeapon {
			selectNextWeapon = true
		}
	}
}
