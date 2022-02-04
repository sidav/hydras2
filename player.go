package main

type player struct {
	dungX, dungY int
	inventory []*item

	// stats
	strength int // how many items can be carried
}

func (p *player) init() {
	p.strength = 5
}
