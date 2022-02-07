package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	_ "github.com/gdamore/tcell/v2"
	"strconv"
)

const (
	dung_x_offset = 0
	dung_y_offset = 1
	roomW, roomH = 3, 3 // not counting walls
)

func (c *consoleIO) renderDungeon(d *dungeon, p *player) {
	c.screen.Clear()

	// dw := len(d.rooms)*(roomW+1)
	dh := len(d.rooms[0])*(roomH+1)

	for rx := range d.rooms {
		for ry := range d.rooms[rx] {
			if d.rooms[rx][ry].wasSeen && d.rooms[rx][ry].isCleared() {
				c.renderRoom(rx, ry, d)
			}
		}
	}
	for rx := range d.rooms {
		for ry := range d.rooms[rx] {
			if d.rooms[rx][ry].wasSeen && !d.rooms[rx][ry].isCleared() {
				c.renderRoom(rx, ry, d)
			}
		}
	}

	c.style = c.style.Foreground(tcell.ColorBlue).Background(tcell.ColorBlack)
	// render player's @
	c.screen.SetCell(p.dungX*(roomW+1)+(roomW+2)/2+dung_x_offset, p.dungY*(roomH+1)+(roomH+2)/2+dung_y_offset, c.style, '@')
	c.renderPlayerDungeonUI(dh+2, d)
	c.screen.Show()
}

func (c *consoleIO) renderRoom(rx, ry int, d *dungeon) {
	cell := d.rooms[rx][ry]
	// render room outline.
	if cell.isCleared() {
		c.style = c.style.Background(tcell.ColorDarkBlue)
	} else {
		c.style = c.style.Background(tcell.ColorDarkRed)
	}
	topLeftX := rx*(roomW+1)+dung_x_offset
	topLeftY := ry*(roomH+1)+dung_y_offset
	runemap := d.layout.CellToCharArray(rx, ry, false, false, false)
	for x := range runemap {
		for y := range runemap[x] {
			runeToDraw := '?'
			switch runemap[x][y] {
			case '#':
				if cell.isCleared() {
					c.setStyle(tcell.ColorBlack, tcell.ColorDarkBlue)
				} else {
					c.setStyle(tcell.ColorBlack, tcell.ColorDarkRed)
				}
				runeToDraw = ' '
			case '1', '2', '3':
				c.setStyle(tcell.ColorBlack, tcell.ColorBlue)
				runeToDraw = '='
			default:
				runeToDraw = runemap[x][y]
				c.resetStyle()
			}
			c.putChar(runeToDraw, x+topLeftX, y+topLeftY)
		}
	}

	enemiesCountStr := strconv.Itoa(len(cell.enemies))
	if enemiesCountStr != "0" {
		c.setStyle(tcell.ColorRed, tcell.ColorBlack)
		c.putString(enemiesCountStr, topLeftX+1, topLeftY+1)
	}
	treasureCountStr := strconv.Itoa(len(cell.treasure))
	if treasureCountStr != "0" {
		c.setStyle(tcell.ColorGreen, tcell.ColorBlack)
		c.putString(treasureCountStr, topLeftX+roomW, topLeftY+1)
	}
}

func (c *consoleIO) renderPlayerDungeonUI(yCoord int, d *dungeon) {
	c.resetStyle()
	var lines = []string{
		fmt.Sprintf("HP: %d/%d ", plr.hitpoints, plr.getMaxHp()),
	}
	if len(plr.keys) > 0 {
		keyLine := "Keys: "
		for k, _ := range plr.keys {
			keyLine += fmt.Sprintf("%d ", k)
		}
		lines = append(lines, keyLine)
	}
	for i := range lines {
		c.putString(lines[i], 0, yCoord+i)
	}
}
