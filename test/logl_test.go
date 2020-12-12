package test

import (
	"bufio"
	"encoding/json"
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
	if level := logl.GetLevel(); level != logl.INFO {
		t.Errorf("Wrong log level %s", level)
	}
	logl.SetLevel(logl.DEBUG)
	if level := logl.GetLevel(); level != logl.DEBUG {
		t.Fatalf("Wrong log level %s", level)
	}
	logl.Info("Test", "ing")

	logl.Debug("Aha!")
	logl.SetLevel(logl.WARN)
	logl.Debug("This should not show")
	logl.Warnf("This should be %d int followed by an error: %s", 1, fmt.Errorf("A new error"))

}

// A better test
func TestFileUnformated(t *testing.T) {

	logl.SetLevel(logl.DEBUG)
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
		if tt, err := time.Parse("2006/01/02 15:04:05", timestamp); err != nil {
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
	if !strings.HasSuffix(line, "This is via log") {
		t.Errorf("log output \"%s\" does not end with \"This is via log\"", line)
	}
}

func TestFileFormated(t *testing.T) {

	logl.SetLevel(logl.DEBUG)
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
		if tt, err := time.Parse("2006/01/02 15:04:05", timestamp); err != nil {
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

func TestJson(t *testing.T) {
	var level logl.Level // = 0
	for level = logl.NONE; level <= logl.TRACE; level++ {
		b, err := json.Marshal(level)
		if err != nil {
			t.Fatalf("Failed to marshal %v", err)
		}
		if fmt.Sprintf("%q", level.String()) != string(b) {
			t.Fatalf("%s != %s", level.String(), string(b))
		}
		var result logl.Level
		err = json.Unmarshal(b, &result)
		if err != nil {
			t.Fatalf("Failed to unmarshal %v", err)
		}
		if result != level {
			t.Fatalf("Wrong result %d (%s)", result, result.String())
		}
		t.Logf("Passed %s", level.String())
	}
}
