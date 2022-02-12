package main

import (
	"github.com/gdamore/tcell/v2"
	"hydras2/text_colors"
)

func (c *consoleIO) getColorByColorTag(tag string) tcell.Color {
	switch tag {
	case "RED":
		return tcell.ColorRed
	case "YELLOW":
		return tcell.ColorYellow
	case "BLUE":
		return tcell.ColorBlue
	case "DARKBLUE":
		return tcell.ColorDarkBlue
	case "CYAN":
		return tcell.ColorLightCyan
	case "DARKCYAN":
		return tcell.ColorDarkCyan
	case "DARKGRAY":
		return tcell.ColorDarkGray
	case "DARKMAGENTA":
		return tcell.ColorDarkMagenta
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
	if !text_colors.IsStringColorTagged(str) {
		c.putUncoloredString(str, x, y)
		return
	}
	offset := 0
	for i := 0; i < len(str); i++ {
		tag := text_colors.GetColorTagNameInStringAtPosition(str, i)
		if tag != "" {
			i += text_colors.COLOR_TAG_LENGTH
			offset += text_colors.COLOR_TAG_LENGTH
			c.setFgColorByColorTag(tag)
			if i >= len(str) {
				break
			}
		}
		c.screen.SetCell(x+i-offset+c.offsetX, y+c.offsetY, c.style, rune(str[i]))
	}
	c.resetStyle()
}

// for word-by-word rendering of text, where only first word has a tag
func (c *consoleIO) putColorTaggedStringNonResetting(str string, x, y int) {
	offset := 0
	for i := 0; i < len(str); i++ {
		tag := text_colors.GetColorTagNameInStringAtPosition(str, i)
		if tag != "" {
			i += text_colors.COLOR_TAG_LENGTH
			offset += text_colors.COLOR_TAG_LENGTH
			c.setFgColorByColorTag(tag)
		}
		c.screen.SetCell(x+i-offset+c.offsetX, y+c.offsetY, c.style, rune(str[i]))
	}
}
