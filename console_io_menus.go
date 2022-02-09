package main

import (
	"github.com/gdamore/tcell"
	"hydras2/text_colors"
)

func (c *consoleIO) showYNSelect(title string, lines []string) bool {
	c.screen.Clear()
	cursor := 0
	longestLineLen := text_colors.TaggedStringLength(title) + 2
	for i := range lines {
		if text_colors.TaggedStringLength(lines[i]) > longestLineLen {
			longestLineLen = text_colors.TaggedStringLength(lines[i])
		}
	}
	for {
		c.setStyle(tcell.ColorBlack, tcell.ColorDarkMagenta)
		c.drawRect(0, 0, longestLineLen+2, len(lines)+2)
		c.resetStyle()
		c.drawStringCenteredAround(title, (longestLineLen+2)/2, 0)
		for i, l := range lines {
			c.putColorTaggedString(l, 1, 1+i)
		}
		if cursor == 0 {
			c.setStyle(tcell.ColorBlack, tcell.ColorWhite)
		} else {
			c.resetStyle()
		}
		c.drawStringCenteredAround("YES", (longestLineLen+2)/3, len(lines)+2)
		if cursor == 1 {
			c.setStyle(tcell.ColorBlack, tcell.ColorWhite)
		} else {
			c.resetStyle()
		}
		c.drawStringCenteredAround("NO", 2*(longestLineLen+2)/3, len(lines)+2)
		c.screen.Show()
		k := c.readKey()
		switch k {
		case "LEFT":
			cursor--
		case "RIGHT":
			cursor++
		case "ENTER":
			return cursor == 0
		case "y":
			return true
		case "n":
			return false
		}
	}
}

func (c *consoleIO) showSelectWindow(title string, lines []string) int {
	c.screen.Clear()
	cursor := 0
	longestLineLen := text_colors.TaggedStringLength(title) + 2
	for i := range lines {
		if len(lines[i]) > longestLineLen {
			longestLineLen = text_colors.TaggedStringLength(lines[i])
		}
	}
	for {
		c.setStyle(tcell.ColorBlack, tcell.ColorDarkMagenta)
		c.drawRect(0, 0, longestLineLen+1, len(lines)+1)
		c.resetStyle()
		c.drawStringCenteredAround(title, (longestLineLen+2)/2, 0)
		for i, l := range lines {
			if i == cursor {
				c.setStyle(tcell.ColorBlack, tcell.ColorWhite)
			} else {
				c.resetStyle()
			}
			c.putColorTaggedString(l, 1, 1+i)
		}
		c.screen.Show()
		k := c.readKey()
		switch k {
		case "UP":
			cursor--
		case "DOWN":
			cursor++
		case "ENTER":
			return cursor
		case "ESCAPE":
			return -1
		}
	}
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
		c.drawRect(0, 0, longestLineLen+1, len(lines)+1)
		c.resetStyle()
		c.drawStringCenteredAround(title, (longestLineLen+2)/2, 0)
		currentPosition := 0
		for i, l := range lines {
			if enabled(i) {
				if i == cursor {
					c.setStyle(tcell.ColorBlack, tcell.ColorWhite)
				} else {
					c.resetStyle()
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

func (c *consoleIO) showInfoWindow(title string, lines []string) {
	longestLineLen := text_colors.TaggedStringLength(title) + 2
	for i := range lines {
		if text_colors.TaggedStringLength(lines[i]) > longestLineLen {
			longestLineLen = text_colors.TaggedStringLength(lines[i])
		}
	}
	for {
		c.setStyle(tcell.ColorBlack, tcell.ColorBlack)
		c.drawFilledRect(' ', 0, 0, longestLineLen+1, len(lines)+2)
		c.setStyle(tcell.ColorBlack, tcell.ColorDarkMagenta)
		c.drawRect(0, 0, longestLineLen+1, len(lines)+2)
		c.resetStyle()
		c.drawStringCenteredAround(title, (longestLineLen+2)/2, 0)
		for i, l := range lines {
			c.putColorTaggedString(l, 1, 1+i)
		}
		c.drawStringCenteredAround("<OK>", (longestLineLen+2)/2, len(lines)+2)
		c.screen.Show()
		k := c.readKey()
		switch k {
		case "ENTER":
			return
		}
	}
}
