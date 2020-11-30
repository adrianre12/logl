package loglevel

import (
	"log"
	golog "log"
	"os"
	"sync"
)

type LogLevel int

const (
	NONE LogLevel = iota
	INFO
	WARN
	DEBUG
)

type Logger struct {
	level LogLevel
	golog *golog.Logger
	S     string
}

var (
	instance *Logger
	once     sync.Once
)

func init() {
	instance = &Logger{level: INFO}
	instance.golog = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func GetLogger() *Logger {
	once.Do(func() {
		instance = &Logger{level: INFO}
		instance.golog = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	})
	return instance
}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *Logger) GetLevel() LogLevel {
	return l.level
}
