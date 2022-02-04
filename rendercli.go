package main

import (
	"github.com/gdamore/tcell"
	_ "github.com/gdamore/tcell/v2"
)

type cliRenderer struct {
	screen                        tcell.Screen
	style                         tcell.Style
	CONSOLE_WIDTH, CONSOLE_HEIGHT int
}

func (c *cliRenderer) init() {
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
	c.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	c.screen.SetStyle(c.style)
	c.CONSOLE_WIDTH, c.CONSOLE_HEIGHT = c.screen.Size()
}

func (c *cliRenderer) close() {
	c.screen.Fini()
}

func (c *cliRenderer) renderDungeon(d *dungeon) {
	chars := *d.layout.WholeMapToCharArray(false, false, false)
	for x := 0; x < len(chars); x++ {
		for y := 0; y < len((chars)[0]); y++ {
			chr := chars[x][y]
			switch chr {
			case '#':
				c.style = c.style.Background(tcell.ColorDarkRed)
				chr = ' '
				break
			default:
				c.style = c.style.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
			}
			c.screen.SetCell(x, y, c.style, chr)
		}
	}
	c.screen.Show()
	c.readKey()
}

func (c *cliRenderer) readKey() {
	for {
		ev := c.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC {
				return
			}
			return
		}
	}
}
