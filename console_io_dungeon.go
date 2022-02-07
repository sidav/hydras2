package main

import (
	"github.com/gdamore/tcell"
	_ "github.com/gdamore/tcell/v2"
	"github.com/sidav/cyclicdungeongenerator/generator/layout_generation"
	"strconv"
)

const (
	dung_x_offset = 0
	dung_y_offset = 1
	roomW, roomH = 3, 3 // not counting walls
)

func (c *consoleIO) renderDungeon(d *dungeon, p *player) {
	c.screen.Clear()

	//dw := len(d.rooms)*(roomW+1)
	//dh := len(d.rooms[0])*(roomH+1)
	//// render outline:
	//c.style = c.style.Background(tcell.ColorDarkBlue)
	//for x := 0; x <= dw; x++ {
	//	c.putChar(' ', x, dung_y_offset)
	//	c.putChar(' ', x, dh+dung_y_offset)
	//}
	//for y := 0; y <= dh; y++ {
	//	c.putChar(' ', 0, y+dung_y_offset)
	//	c.putChar(' ', dw, y+dung_y_offset)
	//}

	for rx := range d.rooms {
		for ry := range d.rooms[rx] {
			if d.rooms[rx][ry].isVisited && d.rooms[rx][ry].isCleared() {
				c.renderRoom(rx, ry, d.layout.GetElement(rx, ry), d.rooms[rx][ry])
			}
		}
	}
	for rx := range d.rooms {
		for ry := range d.rooms[rx] {
			if d.rooms[rx][ry].isVisited && !d.rooms[rx][ry].isCleared() {
				c.renderRoom(rx, ry, d.layout.GetElement(rx, ry), d.rooms[rx][ry])
			}
		}
	}

	c.style = c.style.Foreground(tcell.ColorBlue).Background(tcell.ColorBlack)
	// render player's @
	c.screen.SetCell(p.dungX*(roomW+1)+(roomW+2)/2+dung_x_offset, p.dungY*(roomH+1)+(roomH+2)/2+dung_y_offset, c.style, '@')
	c.screen.Show()
}

func (c *consoleIO) renderRoom(rx, ry int, element *layout_generation.Element, cell *dungeonCell) {
	// render room outline.
	if cell.isCleared() {
		c.style = c.style.Background(tcell.ColorDarkBlue)
	} else {
		c.style = c.style.Background(tcell.ColorDarkRed)
	}
	topLeftX := rx*(roomW+1)+dung_x_offset
	topLeftY := ry*(roomH+1)+dung_y_offset
	for x := topLeftX; x < topLeftX+roomW+1; x++ {
		c.putChar(' ', x, ry*(roomH+1)+dung_y_offset)
		c.putChar(' ', x, (ry+1)*(roomH+1)+dung_y_offset)
	}
	for y := topLeftY; y <= topLeftY+roomH+1; y++ {
		c.putChar(' ', rx*(roomW+1), y)
		c.putChar(' ', (rx+1)*(roomW+1), y)
	}
	// render connections...
	c.resetStyle()
	centerX, centerY := topLeftX+roomW/2+1, topLeftY+roomH/2+1
	doorXOffset := roomW/2+1
	doorYOffset := roomH/2+1
	conns := element.GetAllConnectionsCoords()
	for _, connCoords := range conns {
		conn := element.GetConnectionByCoords(connCoords[0], connCoords[1])
		connChar := ' '
		switch conn.LockNum {
		case 0:
			connChar = '+'
		case 1:
			connChar = '='
		}
		c.putChar(connChar, centerX+doorXOffset*connCoords[0], centerY+doorYOffset*connCoords[1])
	}

	enemiesCountStr := strconv.Itoa(len(cell.enemies))
	if enemiesCountStr != "0" {
		c.setStyle(tcell.ColorRed, tcell.ColorBlack)
		c.putString(enemiesCountStr, topLeftX+1, topLeftY+1)
	}
}
