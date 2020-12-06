package test

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"testing"
	"time"

	"github.com/adrianre12/logl"
)

//Not Much of a test I know, but it proves it works
func TestLogger(t *testing.T) {
	fmt.Println("Start")
	if level := logl.GetLevel(); level != logl.INF {
		t.Errorf("Wrong log level %s", level)
	}
	logl.SetLevel(logl.DBG)
	if level := logl.GetLevel(); level != logl.DBG {
		t.Fatalf("Wrong log level %s", level)
	}
	logl.Info("Test", "ing")

	logl.Debug("Aha!")
	logl.SetLevel(logl.WRN)
	logl.Debug("This should not show")
	logl.Warnf("This should be %d int followed by an error: %s", 1, fmt.Errorf("A new error"))

}

// A better test
func TestFileUnformated(t *testing.T) {

	logl.SetLevel(logl.DBG)
	err := logl.SetFileWriter("tmp/unformated.log", true)
	if err != nil {
		t.Fatal(err)
	}

	logl.Error("Unformated ", 0)
	logl.Warn("Unformated ", 1)
	logl.Info("Unformated ", 2)
	logl.Debug("Unformated ", 3)
	logl.Trace("Unformated ", 4)

	log.Println("This is via log")
	logl.Close() // flush and close the writer

	lf, err := os.Open("tmp/unformated.log")
	if err != nil {
		t.Fatal(err)
	}
	defer lf.Close()

	sc := bufio.NewScanner(lf)
	now := time.Now()
	i := 0
	var l logl.Level

	// test logl writing
	for sc.Scan() && (i < 5) {
		l = logl.Level(i + 1)

		expected := fmt.Sprintf("[%s] Unformated %d", l, i)
		line := sc.Text()
		parts := strings.SplitN(line, " ", 4)
		timestamp := fmt.Sprintf("%s %s", parts[0], parts[1])
		function := (strings.Split(parts[2], ":"))[0]
		msg := parts[3]
		if msg != expected {
			t.Errorf("line %d \"%s\" != \"%s\"", i, msg, expected)
		}

		if function != "logl_test.go" {
			t.Errorf("line %d \"%s\" != \"%s\"", i, function, "logl_test.go")
		}
		if tt, err := time.Parse("2006/01/02 03:04:05", timestamp); err != nil {
			t.Errorf("Line %d, failed to parse timestamp \"%s\"", i, timestamp)
		} else {
			td := now.Sub(tt)

			if math.Abs(td.Seconds()) > 5 {
				t.Errorf("Line %d, timestamp delta > 5s: %f", i, td.Seconds())
			}
		}
		i++
	}
	// test log writing
	line := sc.Text()
	msg := line[len(line)-16:]
	if msg != "This is from log" {
		t.Errorf("log output \"%s\" does not end with \"This is from log\"", msg)
	}
}

func TestFileFormated(t *testing.T) {

	logl.SetLevel(logl.DBG)
	err := logl.SetFileWriter("tmp/formated.log", true)
	if err != nil {
		t.Fatal(err)
	}

	logl.Errorf("Formated %d", 0)
	logl.Warnf("Formated %d", 1)
	logl.Infof("Formated %d", 2)
	logl.Debugf("Formated %d", 3)
	logl.Tracef("Formated %d", 4)

	logl.Close() // flush and close the writer

	lf, err := os.Open("tmp/formated.log")
	if err != nil {
		t.Fatal(err)
	}
	defer lf.Close()

	sc := bufio.NewScanner(lf)
	now := time.Now()
	i := 0
	var l logl.Level

	// test logl writing
	for sc.Scan() && (i < 5) {
		l = logl.Level(i + 1)

		expected := fmt.Sprintf("[%s] Formated %d", l, i)
		line := sc.Text()
		parts := strings.SplitN(line, " ", 4)
		timestamp := fmt.Sprintf("%s %s", parts[0], parts[1])
		function := (strings.Split(parts[2], ":"))[0]
		msg := parts[3]
		if msg != expected {
			t.Errorf("line %d \"%s\" != \"%s\"", i, msg, expected)
		}

		if function != "logl_test.go" {
			t.Errorf("line %d \"%s\" != \"%s\"", i, function, "logl_test.go")
		}
		if tt, err := time.Parse("2006/01/02 03:04:05", timestamp); err != nil {
			t.Errorf("Line %d, failed to parse timestamp \"%s\"", i, timestamp)
		} else {
			td := now.Sub(tt)

			if math.Abs(td.Seconds()) > 5 {
				t.Errorf("Line %d, timestamp delta > 5s: %f", i, td.Seconds())
			}
		}
		i++
	}
}
