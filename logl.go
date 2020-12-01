// Package logl implements logging at levels from NONE to DEBUG,
// The output is filtered by a globally set logging level.
// The objective was to write a logging package that is small and simple.
// The alternatives I looked at were either larger and more complex than my
// applications, or did not do what I required.
// To do this logl wraps an instance of log.Logger and calls the relevant
// functions in log to write the mesages.
package logl

import (
	"bufio"
	"fmt"
	"io"
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

// For conveniance these are copied from the log code.
// These flags define which text to prefix to each log entry generated by the Logger.
// Bits are or'ed together to control what's printed.
// With the exception of the Lmsgprefix flag, there is no
// control over the order they appear (the order listed here)
// or the format they present (as described in the comments).
// The prefix is followed by a colon only when Llongfile or Lshortfile
// is specified.
// For example, flags Ldate | Ltime (or LstdFlags) produce,
//	2009/01/23 01:23:23 message
// while flags Ldate | Ltime | Lmicroseconds | Llongfile produce,
//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

var (
	instance *log.Logger
	level    Level
	file     *os.File
	writer   io.Writer
)

// Called at load time and creates an new instance of log.Logger writing to StdOut and with flags LstdFlags|Lshortfile
func init() {
	instance = log.New(os.Stdout, "", LstdFlags|Lshortfile)
	level = INFO
}

// Flushes and closes the output stream
func Close() {
	if writer != nil {
		if b, ok := writer.(*bufio.Writer); ok {
			b.Flush()
		}
	}
	if file != nil {
		file.Close()
	}
}

// Set the minimum logging level NONE ... DEBUG
func SetLevel(l Level) {
	level = l
}

// The current logging level
func GetLevel() Level {
	return level
}

// Set the io.writer for log.Logger
func SetWriter(out io.Writer) {
	writer = out // so we can handle flushes on close
	instance.SetOutput(out)
}

// Set the io.writer to write to the file specified
func SetFileWriter(fileName string) error {
	var err error
	file, err = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	writer = bufio.NewWriter(file)

	SetWriter(writer)
	return nil
}

// Set the output flags for log.Logger
func SetFlags(flag int) {
	instance.SetFlags(flag)
}

// Logs at the DEBUG level, see Info() for details
func Debug(v ...interface{}) {
	if level < DEBUG {
		return
	}
	log.Print("[DEBUG] ", fmt.Sprint(v...))
}

// Write a log message at the INFO level, but only if the logging level is INFO or higher.
// The message supplied will be prefixed with the logging level
func Info(v ...interface{}) {
	if level < INFO {
		return
	}
	log.Print("[INFO] ", fmt.Sprint(v...))
}

// Logs at the WARN level, see Info() for details
func Warn(v ...interface{}) {
	if level < WARN {
		return
	}
	log.Print("[WARN] ", fmt.Sprint(v...))
}

// Logs at the ERROR level, see Info() for details
func Error(v ...interface{}) {
	if level < ERROR {
		return
	}
	log.Print("[ERROR] ", fmt.Sprint(v...))
}

// Logs at the FATAL level, clss Close() then terminates using os.Exit(1).
// Only use if you really have to!
// See Info() for details
func Fatal(f string, v ...interface{}) {
	log.Print("[FATAL] "+f, fmt.Sprint(v...))
	Close()
	os.Exit(1)
}

// Logs at the DEBUG level, see Info() for details
// Arguments are handled in the manner of fmt.Printf.
func Debugf(f string, v ...interface{}) {
	if level < DEBUG {
		return
	}
	log.Printf("[DEBUG] "+f, fmt.Sprint(v...))
}

// Logs at the INFO level, see Info() for details
// Arguments are handled in the manner of fmt.Printf.
func Infof(f string, v ...interface{}) {
	if level < INFO {
		return
	}
	log.Printf("[INFO] "+f, fmt.Sprint(v...))
}

// Logs at the WARN level, see Info() for details
// Arguments are handled in the manner of fmt.Printf.
func Warnf(f string, v ...interface{}) {
	if level < WARN {
		return
	}
	log.Printf("[WARN] "+f, fmt.Sprint(v...))
}

// Logs at the ERROR level, see Info() for details
// Arguments are handled in the manner of fmt.Printf.
func Errorf(f string, v ...interface{}) {
	if level < ERROR {
		return
	}
	log.Printf("[ERROR] "+f, fmt.Sprint(v...))
}

// Logs at the FATAL level, clss Close() then terminates using os.Exit(1).
// Only use if you really have to!
// Arguments are handled in the manner of fmt.Printf. See Info() for details
func Fatalf(f string, v ...interface{}) {
	log.Printf("[FATAL] "+f, fmt.Sprint(v...))
	Close()
	os.Exit(1)
}