package main

import (
	"github.com/gdamore/tcell"
	"hydras2/entities"
)

func (c *consoleIO) setFgColorByColorTag(tagName string) {
	switch tagName {
	case "RED":
		c.style = c.style.Foreground(tcell.ColorRed)
	case "YELLOW":
		c.style = c.style.Foreground(tcell.ColorYellow)
	case "BLUE":
		c.style = c.style.Foreground(tcell.ColorBlue)
	case "CYAN":
		c.style = c.style.Foreground(tcell.ColorLightCyan)
	case "DARKGRAY":
		c.style = c.style.Foreground(tcell.ColorDarkGray)
	case "RESET":
		c.resetStyle()
	default:
		panic("Y U NO IMPLEMENT")
	}
}

func (c *consoleIO) putColorTaggedString(str string, x, y int) {
	if !entities.IsStringColorTagged(str) {
		c.putUncoloredString(str, x, y)
		return
	}
	offset := 0
	for i := 0; i < len(str); i++ {
		tag := entities.GetColorTagNameInStringAtPosition(str, i)
		if tag != "" {
			i += entities.COLOR_TAG_LENGTH
			offset += entities.COLOR_TAG_LENGTH
			c.setFgColorByColorTag(tag)
		}
		c.screen.SetCell(x+i-offset, y, c.style, rune(str[i]))
	}
	c.resetStyle()
}
