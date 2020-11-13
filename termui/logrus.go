package mywidgets

import (
	"fmt"
	"strings"

	"github.com/gizak/termui/v3/widgets"
	log "github.com/sirupsen/logrus"
)

const (
	maxLogRows = 10
)

// LogrusList An implementation of both a TermUI List Widget and a Logrus Hook
type LogrusList struct {
	widgets.List
	logLevels []log.Level
}

// Levels returns a slice of all currently enable levels to watch for.
// This is part of the Logrus Hook definition
func (lr *LogrusList) Levels() []log.Level {
	return lr.logLevels
}

// Fire Method for recieving new log entries when they are fired.
// This is part of the Logrus Hook definition
func (lr *LogrusList) Fire(entry *log.Entry) error {
	// time := entry.Time
	msg := entry.Message
	lvl := entry.Level
	fields := entry.Data

	strFields := []string{}
	for k, v := range fields {
		strFields = append(strFields, fmt.Sprintf("%s=%v", k, v))
	}

	lr.Rows = append(lr.Rows, fmt.Sprintf("[%s] %s  Fields|%s|", lvl, msg, strings.Join(strFields, ", ")))
	l := len(lr.Rows)
	if l > maxLogRows {
		lr.Rows = lr.Rows[l-maxLogRows : l]
	}
	lr.ScrollBottom()
	return nil
}

// NewLogrusList Creates a new LogrusList to listen for the given log events at the provided levels
func NewLogrusList(lvl ...log.Level) *LogrusList {
	logWidget := LogrusList{
		List:      *widgets.NewList(),
		logLevels: lvl,
	}
	logWidget.Rows = make([]string, 0)

	return &logWidget
}
