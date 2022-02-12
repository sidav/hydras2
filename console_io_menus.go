package main

import (
	"github.com/gdamore/tcell/v2"
	"hydras2/text_colors"
)

func (c *consoleIO) getMenuWindowWidth() int {
	w, _ := c.getConsoleSize()
	if w > 60 {
		return w/2
	}
	if w > 40 {
		return 3*w/4
	}
	return w
}

func (c *consoleIO) showSelectWindow(title string, lines []string) int {
	return c.showSelectWindowWithDisableableOptions(title, lines, func(x int) bool {return true}, true)
}

func (c *consoleIO) showSelectWindowWithDisableableOptions(title string, lines []string,
	enabled func(int) bool, showDisabled bool) int {
	c.screen.Clear()
	cursor := 0
	for i := 0; i < len(lines) && !enabled(cursor); i++ {
		cursor++
	}
	longestLineLen := text_colors.TaggedStringLength(title) + 2
	for i := range lines {
		if len(lines[i]) > longestLineLen {
			longestLineLen = text_colors.TaggedStringLength(lines[i])
		}
	}
	for {
		c.setStyle(tcell.ColorBlack, tcell.ColorDarkMagenta)
		c.drawRect(0, 0, longestLineLen+4, len(lines)+1)
		c.resetStyle()
		c.drawStringCenteredAround(title, (longestLineLen+2)/2, 0)
		currentPosition := 0
		for i, l := range lines {
			if enabled(i) {
				if i == cursor {
					l = "->" + l // c.setStyle(tcell.ColorBlack, tcell.ColorWhite)
				} else {
					l = l+"  "
				}
				c.putColorTaggedString(l, 1, 1+currentPosition)
				currentPosition++
			} else if showDisabled {
				c.setStyle(tcell.ColorDarkGray, tcell.ColorBlack)
				c.putColorTaggedString(l, 1, 1+currentPosition)
				currentPosition++
			}
		}
		c.screen.Show()
		k := c.readKey()
		switch k {
		case "UP":
			for i := 0; i == 0 || i < len(lines) && !enabled(cursor); i++ {
				cursor--
				if cursor < 0 {
					cursor = len(lines) - 1
				}
			}
		case "DOWN":
			for i := 0; i == 0 || i < len(lines) && !enabled(cursor); i++ {
				cursor++
				if cursor >= len(lines) {
					cursor = 0
				}
			}
		case "ENTER":
			return cursor
		case "ESCAPE":
			return -1
		}
	}
}

func (c *consoleIO) showYNSelect(title, text string) bool {
	windowWidth := c.getMenuWindowWidth()
	c.screen.Clear()
	cursor := 0
	for {
		c.resetStyle()
		windowHeight := c.putWrappedTextInRect(text, 1, 1, windowWidth-1)+1
		c.setStyle(tcell.ColorBlack, tcell.ColorDarkMagenta)
		c.drawRect(0, 0, windowWidth, windowHeight)
		c.setStyle(tcell.ColorDarkMagenta, tcell.ColorBlack)
		c.drawStringCenteredAround(" " + title + " ", (windowWidth+2)/2, 0)
		if cursor == 0 {
			c.setStyle(tcell.ColorBlack, tcell.ColorWhite)
		} else {
			c.resetStyle()
		}
		c.drawStringCenteredAround("< YES", (windowWidth+2)/3, windowHeight)
		if cursor == 1 {
			c.setStyle(tcell.ColorBlack, tcell.ColorWhite)
		} else {
			c.resetStyle()
		}
		c.drawStringCenteredAround("NO >", 2*(windowWidth+2)/3, windowHeight)
		c.screen.Show()
		k := c.readKey()
		switch k {
		case "LEFT":
			cursor=0
		case "RIGHT":
			cursor=1
		case "ENTER":
			return cursor == 0
		case "y":
			return true
		case "n":
			return false
		}
	}
}

func (c *consoleIO) showInfoWindow(title, text string) {
	windowWidth := c.getMenuWindowWidth()
	for {
		c.resetStyle()
		textHeight := c.putWrappedTextInRect(text, 1, 1, windowWidth-1)
		c.setStyle(tcell.ColorBlack, tcell.ColorDarkMagenta)
		c.drawRect(0, 0, windowWidth+1, textHeight+1)
		c.setStyle(tcell.ColorDarkMagenta, tcell.ColorBlack)
		c.drawStringCenteredAround(" "+title+" ", (windowWidth+2)/2, 0)
		c.setStyle(tcell.ColorBlack, tcell.ColorWhite)
		c.drawStringCenteredAround("<OK>", (windowWidth+2)/2, textHeight+1)
		c.screen.Show()
		k := c.readKey()
		switch k {
		case "ENTER":
			return
		}
	}
}
