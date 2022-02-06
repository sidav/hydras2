package main

type player struct {
	dungX, dungY int
	inventory []*item

	// stats
	strength int // how many items can be carried

	// battlefield-only vars
	x, y int
}

func (p *player) init() {
	p.strength = 5
}
