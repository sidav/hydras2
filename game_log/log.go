package game_log

import (
	"fmt"
	"strings"
)

type logMessage struct {
	Message string
	count   int
}

func (m *logMessage) GetText() string {
	if m.count > 1 {
		return fmt.Sprintf("%s (x%d)", m.Message, m.count)
	} else {
		return m.Message
	}
}

type GameLog struct {
	LastMessages  []logMessage
	logWasChanged bool
}

func (l *GameLog) Init(length int) {
	l.LastMessages = make([]logMessage, length)
}

func (l *GameLog) Clear() {
	l.LastMessages = make([]logMessage, len(l.LastMessages))
	//for i := range l.LastMessages {
	//	l.LastMessages[i].count = 0
	//	l.LastMessages[i].Message = ""
	//}
}

func (l *GameLog) AppendMessage(msg string) {
	msg = capitalize(msg)
	if l.LastMessages[len(l.LastMessages)-1].Message == msg {
		l.LastMessages[len(l.LastMessages)-1].count++
	} else {
		for i := 0; i < len(l.LastMessages)-1; i++ {
			l.LastMessages[i] = l.LastMessages[i+1]
		}
		l.LastMessages[len(l.LastMessages)-1] = logMessage{Message: msg, count:1}
	}
	l.logWasChanged = true
}

func (l *GameLog) AppendMessagef(msg string, zomg ...interface{}) {
	msg = fmt.Sprintf(msg, zomg...)
	l.AppendMessage(msg)
}

//func (l *GameLog) Warning(msg string) {
//	l.AppendMessage(msg)
//	l.Last_msgs[len(l.Last_msgs) - 1].color = cw.YELLOW
//}
//
//func (l *GameLog) Warningf(msg string, zomg interface{}) {
//	l.AppendMessagef(msg, zomg)
//	l.Last_msgs[len(l.Last_msgs) - 1].color = cw.YELLOW
//}

func (l *GameLog) WasChanged() bool {
	was := l.logWasChanged
	l.logWasChanged = false
	return was
}

func capitalize(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(string(s[0])) + s[1:]
	}
	return s 
}
