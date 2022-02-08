package main

import (
	"github.com/gdamore/tcell"
	_ "github.com/gdamore/tcell/v2"
	"hydras2/entities"
)

type consoleIO struct {
	screen                        tcell.Screen
	style                         tcell.Style
	CONSOLE_WIDTH, CONSOLE_HEIGHT int
}

func (c *consoleIO) init() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	var e error
	c.screen, e = tcell.NewScreen()
	if e != nil {
		panic(e)
	}
	if e = c.screen.Init(); e != nil {
		panic(e)
	}
	// c.screen.EnableMouse()
	c.setStyle(tcell.ColorWhite, tcell.ColorBlack)
	c.screen.SetStyle(c.style)
	c.CONSOLE_WIDTH, c.CONSOLE_HEIGHT = c.screen.Size()
}

func (c *consoleIO) close() {
	c.screen.Fini()
}

func (c *consoleIO) readKey() string {
	for {
		ev := c.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC {
				return "EXIT"
			}
			return eventToKeyString(ev)
		}
	}
}

func (c *consoleIO) showYNSelect(title string, lines []string) bool {
	c.screen.Clear()
	cursor := 0
	longestLineLen := entities.TaggedStringLength(title)+2
	for i := range lines {
		if entities.TaggedStringLength(lines[i]) > longestLineLen {
			longestLineLen = entities.TaggedStringLength(lines[i])
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
	longestLineLen := entities.TaggedStringLength(title)+2
	for i := range lines {
		if len(lines[i]) > longestLineLen {
			longestLineLen = entities.TaggedStringLength(lines[i])
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
		enabled func(int)bool, showDisabled bool) int {
	c.screen.Clear()
	cursor := 0
	for i := 0; i < len(lines) && !enabled(cursor); i++ {
		cursor++
	}
	longestLineLen := entities.TaggedStringLength(title)+2
	for i := range lines {
		if len(lines[i]) > longestLineLen {
			longestLineLen = entities.TaggedStringLength(lines[i])
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
					cursor = len(lines)-1
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
	longestLineLen := entities.TaggedStringLength(title)+2
	for i := range lines {
		if entities.TaggedStringLength(lines[i]) > longestLineLen {
			longestLineLen = entities.TaggedStringLength(lines[i])
		}
	}
	for {
		c.setStyle(tcell.ColorBlack, tcell.ColorBlack)
		c.drawFilledRect(' ',0, 0, longestLineLen+1, len(lines)+2)
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

func eventToKeyString(ev *tcell.EventKey) string {
	switch ev.Key() {
	case tcell.KeyUp:
		return "UP"
	case tcell.KeyRight:
		return "RIGHT"
	case tcell.KeyDown:
		return "DOWN"
	case tcell.KeyLeft:
		return "LEFT"
	case tcell.KeyEscape:
		return "ESCAPE"
	case tcell.KeyEnter:
		return "ENTER"
	case tcell.KeyBackspace, tcell.KeyBackspace2:
		return "BACKSPACE"
	case tcell.KeyTab:
		return "TAB"
	case tcell.KeyDelete:
		return "DELETE"
	case tcell.KeyInsert:
		return "INSERT"
	case tcell.KeyEnd:
		return "END"
	case tcell.KeyHome:
		return "HOME"
	default:
		return string(ev.Rune())
	}
}

func (c *consoleIO) putChar(chr rune, x, y int) {
	c.screen.SetCell(x, y, c.style, chr)
}

func (c *consoleIO) putUncoloredString(str string, x, y int) {
	for i := 0; i < len(str); i++ {
		c.screen.SetCell(x+i, y, c.style, rune(str[i]))
	}
}

func (c *consoleIO) setStyle(fg, bg tcell.Color) {
	c.style = c.style.Background(bg).Foreground(fg)
}

func (c *consoleIO) resetStyle() {
	c.setStyle(tcell.ColorWhite, tcell.ColorBlack)
}

func (c *consoleIO) drawFilledRect(char rune, fx, fy, w, h int) {
	for x := fx; x <= fx+w; x++ {
		for y := fy; y <= fy+h; y++ {
			c.putChar(char, x, y)
		}
	}
}

func (c *consoleIO) drawRect(fx, fy, w, h int) {
	for x := fx; x <= fx+w; x++ {
		c.putChar(' ', x, fy)
		c.putChar(' ', x, fy+h)
	}
	for y := fy; y <= fy+h; y++ {
		c.putChar(' ', fx, y)
		c.putChar(' ', fx+w, y)
	}
}

func (c *consoleIO) drawStringCenteredAround(s string, x, y int) {
	c.putColorTaggedString(s, x-entities.TaggedStringLength(s)/2, y)
}
