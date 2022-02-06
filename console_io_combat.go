package main

import (
	"github.com/gdamore/tcell"
	"strconv"
)

const (
	bf_x_offset = 1
	bf_y_offset = 2
)

func (c *consoleIO) renderBattlefield(b *battlefield) {
	c.screen.Clear()
	c.putString("COMBAT: ", 0, 0)
	bfW, bfH := len(b.tiles), len(b.tiles[0])

	// render outline:
	c.style = c.style.Background(tcell.ColorDarkRed)
	for i := 0; i <= bfW+1; i++ {
		c.putChar(' ', 0, i+1)
		c.putChar(' ', bfW+1, i+1)
		c.putChar(' ', i, 1)
		c.putChar(' ', i, bfH+2)
	}
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
	c.screen.Show()
}

func (c *consoleIO) renderEnemy(e *enemy) {
	strForHeads := "?"
	if e.heads < 10 {
		strForHeads = strconv.Itoa(e.heads)
	}
	c.style = c.style.Foreground(tcell.ColorRed).Background(tcell.ColorBlack)
	c.putString(strForHeads, e.x+bf_x_offset, e.y+bf_y_offset)
}
