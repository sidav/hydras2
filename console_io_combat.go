package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"strconv"
)

const (
	bf_x_offset = 1
	bf_y_offset = 2
)

func (c *consoleIO) renderBattlefield(b *battlefield) {
	c.screen.Clear()
	c.putColorTaggedString("COMBAT: ", 0, 0)
	bfW, bfH := len(b.tiles), len(b.tiles[0])

	// render outline:
	c.setStyle(tcell.ColorWhite, tcell.ColorDarkRed)
	c.drawRect(0, 1, bfW+1, bfH+1)
	// render the battlefield itself
	for x := range b.tiles {
		for y := range b.tiles[x] {
			switch b.tiles[x][y] {
			case TILE_WALL:
				c.style = c.style.Background(tcell.ColorDarkRed)
				c.putChar(' ', x+bf_x_offset, y+bf_y_offset)
			case TILE_FLOOR:
				c.resetStyle()
				c.putChar('.', x+bf_x_offset, y+bf_y_offset)
			}
		}
	}
	for i := range b.enemies {
		c.renderEnemy(b.enemies[i])
	}
	c.resetStyle()
	c.putChar('@', b.player.x+bf_x_offset, b.player.y+bf_y_offset)
	c.renderPlayerBattlefieldUI(bf_x_offset+bfW+1, b)
	c.screen.Show()
}

func (c *consoleIO) renderEnemy(e *enemy) {
	strForHeads := c.getCharForEnemy(e.heads)
	c.style = c.style.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)
	c.putChar(strForHeads, e.x+bf_x_offset, e.y+bf_y_offset)
}

func (c *consoleIO) renderPlayerBattlefieldUI(xCoord int, b *battlefield) {
	var lines = []string{
		fmt.Sprintf("HP: %d/%d", b.player.hitpoints, b.player.getMaxHp()),
		fmt.Sprintf("1) Wpn: %s", b.player.currentWeapon.GetName()),
		fmt.Sprintf("2) Itm: %dx %s",
			b.player.currentConsumable.AsConsumable.Amount,
			b.player.currentConsumable.GetName()),
		"ENEMIES:",
	}
	for i := range b.enemies {
		lines = append(lines, fmt.Sprintf("%s %s", string(c.getCharForEnemy(b.enemies[i].heads)), b.enemies[i].getName()))
	}
	for i := range lines {
		c.putColorTaggedString(lines[i], xCoord, i+1)
	}
}

func (c *consoleIO) getCharForEnemy(heads int) rune {
	if heads < 10 {
		return rune(strconv.Itoa(heads)[0])
	} else if heads < 16 {
		return []rune{'A', 'B', 'C', 'D', 'E', 'F'}[heads-10]
	}
	return '?'
}
