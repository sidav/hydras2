package main

import (
	"github.com/gdamore/tcell"
	"hydras2/entities"
)

func (c *consoleIO) putColorTaggedString(str string, x, y int) {
	if !entities.IsStringColorTagged(str) {
		c.putUncoloredString(str, x, y)
		return
	}
	offset := 0
	for i := 0; i < len(str); i++ {
		skip := true
		switch entities.GetColorTagNameInStringAtPosition(str, i) {
		case "RED":
			c.setStyle(tcell.ColorRed, tcell.ColorBlack)
		case "BLUE":
			c.setStyle(tcell.ColorBlue, tcell.ColorBlack)
		case "CYAN":
			c.setStyle(tcell.ColorDarkCyan, tcell.ColorBlack)
		case "DARKGRAY":
			c.setStyle(tcell.ColorDarkGray, tcell.ColorBlack)
		case "RESET":
			c.resetStyle()
		default:
			skip = false
		}
		if skip {
			i += entities.COLOR_TAG_LENGTH
			offset += entities.COLOR_TAG_LENGTH
		}
		c.screen.SetCell(x+i-offset, y, c.style, rune(str[i]))
	}
	c.resetStyle()
}
