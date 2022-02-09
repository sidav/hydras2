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
	for _, e := range b.enemies {
		c.renderEnemyAtCoords(e, b.currentTick, e.x+bf_x_offset, e.y+bf_y_offset)
	}
	c.resetStyle()
	c.putChar('@', b.player.x+bf_x_offset, b.player.y+bf_y_offset)
	c.renderPlayerBattlefieldUI(bf_x_offset+bfW+1, b)
	c.renderLogAt(log, 0, bf_y_offset)
	c.screen.Show()
}

func (c *consoleIO) renderEnemyAtCoords(e *enemy, tick, x, y int) {
	strForHeads := c.getCharForEnemy(e.heads)
	colorTags := e.element.GetColorTags()
	switch len(colorTags) {
	case 0:
		c.resetStyle()
	case 1:
		c.setFgColorByColorTag(colorTags[0])
	default:
		// magic number 3 is just randomly chosen small prime. 5 or 7 or 11 would also work.
		c.setFgColorByColorTag(colorTags[tick/3%len(colorTags)])
	}
	c.putChar(strForHeads, x, y)
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
	enemiesLinesStart := len(lines)
	for i := range b.enemies {
		lines = append(lines, fmt.Sprintf("  %s (%s)",
			b.enemies[i].getName(),
			getAttackDescriptionString(b.player.currentWeapon, b.enemies[i]),
		))
	}
	for i := range lines {
		c.putColorTaggedString(lines[i], xCoord, i+1)
	}
	// render enemies for those enemy lines
	for i, e := range b.enemies {
		c.renderEnemyAtCoords(e, b.currentTick, xCoord, enemiesLinesStart+i+1)
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
