package logger

import (
	"fmt"
	"time"
)

type (
	Logger struct {
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
	Panic
)

func (l logLevel) String() string {
	return []string{
		"Info",
		"Notice",
		"Debug",
		"Warning",
		"Error",
		"Critical",
		"Panic",
	}[l]
}

func (l logLevel) getFormat() string {
	return []string{
		"\033[1;34m[INFO]\033[0m 	%s\n",
		"\033[1;36m[NOTICE]\033[0m 	%s\n",
		"\033[1;36m[DEBUG]\033[0m 	%s\n",
		"\033[1;33m[WARNING]\033[0m %s\n",
		"\033[1;31m[ERROR]\033[0m 	%s\n",
		"\033[0;31m[CRITIC]\033[0m 	%s\n",
		"\033[0;31m[CRITIC]\033[0m 	%s\n",
	}[l]
}

func New(debugMode bool) Logger {
	return Logger{
		debugMode: debugMode,
	}
}

func (l *Logger) Use(f func(*Log)) {
	l.middles = append(l.middles, f)
}

func (l *Logger) Log(log *Log) {
	for _, m := range l.middles {
		m(log)
	}
	if !l.debugMode && log.Level == Debug {
		return
	}

	fmt.Printf(log.Level.getFormat(), log.Message)
	if log.Level == Panic {
		panic(log.Message)
	}
}

func (l *Logger) logf(level logLevel, format string, a ...any) {
	l.Log(&Log{
		Level:   level,
		Message: fmt.Sprintf(format, a...),
	})
}

func (l *Logger) log(level logLevel, a ...any) {
	l.Log(&Log{
		Level:   level,
		Message: fmt.Sprintf("%s", fmt.Sprint(a...)),
	})
}

func (l *Logger) Infof(format string, a ...any) {
	l.logf(Info, format, a...)
}
func (l *Logger) Noticef(format string, a ...any) {
	l.logf(Notice, format, a...)
}
func (l *Logger) Debugf(format string, a ...any) {
	l.logf(Debug, format, a...)
}
func (l *Logger) Warningf(format string, a ...any) {
	l.logf(Warning, format, a...)
}
func (l *Logger) Errorf(format string, a ...any) {
	l.logf(Error, format, a...)
}
func (l *Logger) Criticalf(format string, a ...any) {
	l.logf(Critical, format, a...)
}
func (l *Logger) Panicf(format string, a ...any) {
	l.logf(Panic, format, a...)
}

func (l *Logger) Info(a ...any) {
	l.log(Info, fmt.Sprint(a...))
}
func (l *Logger) Notice(a ...any) {
	l.log(Notice, fmt.Sprint(a...))
}
func (l *Logger) Debug(a ...any) {
	l.log(Debug, fmt.Sprint(a...))
}
func (l *Logger) Warning(a ...any) {
	l.log(Warning, fmt.Sprint(a...))
}
func (l *Logger) Error(a ...any) {
	l.log(Error, fmt.Sprint(a...))
}
func (l *Logger) Critical(a ...any) {
	l.log(Critical, fmt.Sprint(a...))
}
func (l *Logger) Panic(a ...any) {
	l.log(Panic, fmt.Sprint(a...))
}

func Timestamper(log *Log) {
	log.Message = fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), log.Message)
}
