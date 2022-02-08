package main

func (b *battlefield) startCombatLoop() {
	for !b.battleEnded {
		io.renderBattlefield(b)
		b.workPlayerInput()
		for _, e := range b.enemies {
			b.actAsEnemy(e)
		}
		b.endTurnCleanup()
		b.currentTick++
	}
}
