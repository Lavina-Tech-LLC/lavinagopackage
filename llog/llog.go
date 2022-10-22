package llog

import (
	"fmt"
)

type (
	logger struct {
		middles   []func(*Log)
		debugMode bool
	}
	Log struct {
		Level   logLevel
		Message string
	}
	logLevel int
)

const (
	Info logLevel = iota
	Notice
	Debug
	Warning
	Error
	Critical
)

func (l logLevel) String() string {
	return []string{
		"Info",
		"Notice",
		"Debug",
		"Warning",
		"Error",
		"Critical",
	}[l]
}

func (l logLevel) getFormat() string {
	return []string{
		"\033[1;34m[INFO]\033[0m %s\n",
		"\033[1;36m[NOTICE]\033[0m %s\n",
		"\033[1;36m[DEBUG]\033[0m %s\n",
		"\033[1;33m[WARNING]\033[0m %s\n",
		"\033[1;31m[ERROR]\033[0m %s\n",
		"\033[0;31m[CRITICAL]\033[0m %s\n",
	}[l]
}

func New(debugMode bool) logger {
	return logger{
		debugMode: debugMode,
	}
}

func (l *logger) Use(f func(*Log)) {
	l.middles = append(l.middles, f)
}

func (l *logger) Log(log *Log) {
	for _, m := range l.middles {
		m(log)
	}
	if !l.debugMode && log.Level == Debug {
		return
	}

	fmt.Printf(log.Level.getFormat(), log.Message)
}

func (l *logger) logf(level logLevel, format string, a ...any) {
	l.Log(&Log{
		Level:   level,
		Message: fmt.Sprintf(format, a...),
	})
}

func (l *logger) Infof(format string, a ...any) {
	l.logf(Info, format, a...)
}
func (l *logger) Noticef(format string, a ...any) {
	l.logf(Notice, format, a...)
}
func (l *logger) Debugf(format string, a ...any) {
	l.logf(Debug, format, a...)
}
func (l *logger) Warningf(format string, a ...any) {
	l.logf(Warning, format, a...)
}
func (l *logger) Errorf(format string, a ...any) {
	l.logf(Error, format, a...)
}
func (l *logger) Criticalf(format string, a ...any) {
	l.logf(Critical, format, a...)
}
