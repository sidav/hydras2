package main

import (
	"github.com/gdamore/tcell"
	"hydras2/entities"
)

func (c *consoleIO) getColorByColorTag(tag string) tcell.Color {
	switch tag {
	case "RED":
		return tcell.ColorRed
	case "YELLOW":
		return tcell.ColorYellow
	case "BLUE":
		return tcell.ColorBlue
	case "CYAN":
		return tcell.ColorLightCyan
	case "DARKGRAY":
		return tcell.ColorDarkGray
	default:
		panic("Y U NO IMPLEMENT")
	}
} 

func (c *consoleIO) setFgColorByColorTag(tagName string) {
	if tagName == "RESET" {
		c.resetStyle()
	} else {
		c.style = c.style.Foreground(c.getColorByColorTag(tagName))
	}
}

func (c *consoleIO) setBgColorByColorTag(tagName string) {
	if tagName == "RESET" {
		c.resetStyle()
	} else {
		c.style = c.style.Background(c.getColorByColorTag(tagName))
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
		c.screen.SetCell(x+i-offset+c.offsetX, y+c.offsetY, c.style, rune(str[i]))
	}
	c.resetStyle()
}
