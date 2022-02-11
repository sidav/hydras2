package main

import (
	"github.com/gdamore/tcell"
	_ "github.com/gdamore/tcell/v2"
	"hydras2/game_log"
	"hydras2/text_colors"
	"strings"
)

type consoleIO struct {
	screen                        tcell.Screen
	style                         tcell.Style
	CONSOLE_WIDTH, CONSOLE_HEIGHT int
	offsetX, offsetY              int
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

func (c *consoleIO) setOffsets(x, y int) {
	c.offsetX = x
	c.offsetY = y
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
	c.screen.SetCell(x+c.offsetX, y+c.offsetY, c.style, chr)
}

func (c *consoleIO) putUncoloredString(str string, x, y int) {
	for i := 0; i < len(str); i++ {
		c.screen.SetCell(x+i+c.offsetX, y+c.offsetY, c.style, rune(str[i]))
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
	c.putColorTaggedString(s, x-text_colors.TaggedStringLength(s)/2, y)
}


func (c *consoleIO) renderLogAt(log *game_log.GameLog, x, y int) {
	c.setOffsets(x, y)
	for i, m := range log.LastMessages {
		c.resetStyle()
		c.putColorTaggedString(m.GetText(), 0, i)
	}
	c.setOffsets(0, 0)
}

func (c *consoleIO) putWrappedTextInRect(text string, x, y, w int) {
	const newLineOffset = 2
	currentLine := 0
	currentLineLength := newLineOffset
	textSplitByLines := strings.Split(text, "\n")
	for _, t := range textSplitByLines {
		lineSplitByWords := strings.Split(t, " ")
		for _, word := range lineSplitByWords {
			if currentLineLength+len(word) > w {
				currentLine++
				currentLineLength = 0
			}
			c.putColorTaggedString(word+" ", x+currentLineLength, y+currentLine)
			currentLineLength += len(word)+1
		}
		currentLine++
		currentLineLength = newLineOffset
	}
}
