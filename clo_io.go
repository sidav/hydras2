package main

import (
	"github.com/gdamore/tcell"
	_ "github.com/gdamore/tcell/v2"
)

type cliIO struct {
	screen                        tcell.Screen
	style                         tcell.Style
	CONSOLE_WIDTH, CONSOLE_HEIGHT int
}

func (c *cliIO) init() {
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

func (c *cliIO) close() {
	c.screen.Fini()
}

func (c *cliIO) renderDungeon(d *dungeon, p *player) {
	c.screen.Clear()
	chars := *d.layout.WholeMapToCharArray(false, false, false)
	for x := 0; x < len(chars); x++ {
		for y := 0; y < len((chars)[0]); y++ {
			chr := chars[x][y]
			lx, ly := x/5, y/5 // coords of dungeonCell IN LAYOUT
			if !d.rooms[lx][ly].isVisited {
				continue
			}
			switch chr {
			case '#':
				if d.rooms[lx][ly].isCleared() {
					c.style = c.style.Background(tcell.ColorDarkBlue)
				} else {
					c.style = c.style.Background(tcell.ColorDarkRed)
				}
				chr = ' '
				break
			default:
				c.style = c.style.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
			}
			c.putChar(chr, x, y)
		}
	}
	c.style = c.style.Foreground(tcell.ColorBlue).Background(tcell.ColorBlack)
	c.screen.SetCell(p.dungX*5+2, p.dungY*5+2, c.style, '@')
	c.screen.Show()
}

func (c *cliIO) readKey() string {
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


func (c *cliIO) showSelectWindow(title string, lines []string) int {
	cursor := 0
	for {
		c.putString(title, 1, 0)
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

func (c *cliIO) putChar(chr rune, x, y int) {
	c.screen.SetCell(x, y, c.style, chr)
}

func (c *cliIO) putString(str string, x, y int) {
	for i := 0; i < len(str); i++ {
		c.screen.SetCell(x+i, y, c.style, rune(str[i]))
	}
}

