package main

type player struct {
	dungX, dungY  int
	currentWeapon *item
	currentItem   *item
	inventory     []*item
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

	p.inventory = append(p.inventory, &item{
		asConsumable: nil,
		asWeapon: &itemWeapon{
			weaponType: WTYPE_SUBSTRACTOR,
			damage:     2,
		},
	})
	p.inventory = append(p.inventory, &item{
		asConsumable: nil,
		asWeapon: &itemWeapon{
			weaponType: WTYPE_SUBSTRACTOR,
			damage:     1,
		},
	})
	p.selectNextWeapon()
}

func (p *player) getMaxHp() int {
	return p.vitality/2
}

func (p *player) selectNextWeapon() {
	// shitty code ahead
	selectNextWeapon := p.currentWeapon == nil
	for i := 0; ; i++ {
		if p.inventory[i].asWeapon != nil && selectNextWeapon {
			p.currentWeapon = p.inventory[i]
			return
		}
		if p.inventory[i] == p.currentWeapon {
			selectNextWeapon = true
		}
	}
}
