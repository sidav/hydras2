package main

import (
	"github.com/gdamore/tcell"
	_ "github.com/gdamore/tcell/v2"
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
	longestLineLen := len(title)
	for i := range lines {
		if len(lines[i]) > longestLineLen {
			longestLineLen = len(lines[i])
		}
	}
	for {
		c.setStyle(tcell.ColorBlack, tcell.ColorDarkMagenta)
		c.drawRect(0, 0, longestLineLen+2, len(lines)+2)
		c.resetStyle()
		c.drawStringCenteredAround(title, (longestLineLen+2)/2, 0)
		for i, l := range lines {
			c.putString(l, 1, 1+i)
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
	cursor := 0
	longestLineLen := len(title)
	for i := range lines {
		if len(lines[i]) > longestLineLen {
			longestLineLen = len(lines[i])
		}
	}
	for {
		c.setStyle(tcell.ColorBlack, tcell.ColorDarkMagenta)
		c.drawRect(0, 0, longestLineLen, len(lines)+1)
		c.resetStyle()
		c.drawStringCenteredAround(title, (longestLineLen+2)/2, 0)
		for i, l := range lines {
			if i == cursor {
				l = "> " + l
			}
			c.putString(l+ "  ", 0, 1+i)
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
		}
	}
}

func (c *consoleIO) showInfoWindow(title string, lines []string) {
	longestLineLen := len(title)
	for i := range lines {
		if len(lines[i]) > longestLineLen {
			longestLineLen = len(lines[i])
		}
	}
	for {
		c.setStyle(tcell.ColorBlack, tcell.ColorBlack)
		c.drawFilledRect(0, 0, longestLineLen+1, len(lines)+2)
		c.setStyle(tcell.ColorBlack, tcell.ColorDarkMagenta)
		c.drawRect(0, 0, longestLineLen+1, len(lines)+2)
		c.resetStyle()
		c.drawStringCenteredAround(title, (longestLineLen+2)/2, 0)
		for i, l := range lines {
			c.putString(l, 1, 1+i)
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

func (c *consoleIO) setStyle(fg, bg tcell.Color) {
	c.style = c.style.Background(bg).Foreground(fg)
}

func (c *consoleIO) resetStyle() {
	c.setStyle(tcell.ColorWhite, tcell.ColorBlack)
}

func (c *consoleIO) putString(str string, x, y int) {
	for i := 0; i < len(str); i++ {
		c.screen.SetCell(x+i, y, c.style, rune(str[i]))
	}
}

func (c *consoleIO) drawFilledRect(fx, fy, w, h int) {
	for x := fx; x <= fx+w; x++ {
		for y := fy; y <= fy+h; y++ {
			c.putChar(' ', x, y)
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
	c.putString(s, x-len(s)/2, y)
}
