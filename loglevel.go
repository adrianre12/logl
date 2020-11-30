package loglevel

import (
	"fmt"
	"log"
	"os"
)

type Level int

const (
	NONE Level = iota
	ERROR
	WARN
	INFO
	DEBUG
)

func (l Level) String() string {
	return []string{"NONE", "ERROR", "WARN", "INFO", "DEBUG"}[l]
}

var (
	instance *log.Logger
	level    Level
)

func init() {
	instance = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	level = INFO
}

func SetLevel(l Level) {
	level = l
}

func GetLevel() Level {
	return level
}

func Debug(v ...interface{}) {
	if level < DEBUG {
		return
	}
	log.Print("[DEBUG] ", fmt.Sprint(v...))
}

func Info(v ...interface{}) {
	if level < INFO {
		return
	}
	log.Print("[INFO] ", fmt.Sprint(v...))
}

func Warn(v ...interface{}) {
	if level < WARN {
		return
	}
	log.Print("[WARN] ", fmt.Sprint(v...))
}

func Error(v ...interface{}) {
	if level < ERROR {
		return
	}
	log.Print("[ERROR] ", fmt.Sprint(v...))
}

func Fatal(f string, v ...interface{}) {
	log.Fatal("[FATAL] "+f, fmt.Sprint(v...))
}

func Debugf(f string, v ...interface{}) {
	if level < DEBUG {
		return
	}
	log.Printf("[DEBUG] "+f, fmt.Sprint(v...))
}

func Infof(f string, v ...interface{}) {
	if level < INFO {
		return
	}
	log.Printf("[INFO] "+f, fmt.Sprint(v...))
}

func Warnf(f string, v ...interface{}) {
	if level < WARN {
		return
	}
	log.Printf("[WARN] "+f, fmt.Sprint(v...))
}

func Errorf(f string, v ...interface{}) {
	if level < ERROR {
		return
	}
	log.Printf("[ERROR] "+f, fmt.Sprint(v...))
}

func Fatalf(f string, v ...interface{}) {
	log.Fatalf("[FATAL] "+f, fmt.Sprint(v...))
}
