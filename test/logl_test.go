package test

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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
func TestFileWrite(t *testing.T) {
	logl.SetLevel(logl.DBG)
	err := logl.SetFileWriter("tmp/test.log", true)
	if err != nil {
		t.Fatal(err)
	}

	logl.Error("Line ", 1)
	logl.Warn("Line ", 2)
	logl.Info("Line ", 3)
	logl.Debug("Line ", 4)
	logl.Errorf("Line %d", 5)
	logl.Warnf("Line %d", 6)
	logl.Infof("Line %d", 7)
	logl.Debugf("Line %d", 8)

	log.Println("This is from log")
	logl.Close()

	lf, err := os.Open("tmp/test.log")
	if err != nil {
		t.Fatal(err)
	}

	defer lf.Close()

	sc := bufio.NewScanner(lf)

	now := time.Now()
	i := 0
	var l logl.Level

	// test logl writing
	for sc.Scan() && (i < 8) {
		l = logl.Level(math.Mod(float64(i), 4.0) + 1)
		i++

		line := sc.Text()
		expected := fmt.Sprintf("[%s] Line %d", l, i)
		msg := line[len(line)-12:]
		if msg != expected {
			t.Errorf("line %d \"%s\" != \"%s\"", i, msg, expected)
		}
		function := line[20 : len(line)-17]
		if function != "logl_test.go" {
			t.Errorf("line %d \"%s\" != \"%s\"", i, function, "logl_test.go")
		}
		if tt, err := time.Parse("2006/01/02 03:04:05", line[:19]); err != nil {
			t.Errorf("Line %d, failed to parse timestamp \"%s\"", i, line[:20])
		} else {
			td := now.Sub(tt)

			if math.Abs(td.Seconds()) > 5 {
				t.Errorf("Line %d, timestamp delta > 5s: %f", i, td.Seconds())
			}
		}
	}
	// test log writing
	line := sc.Text()
	msg := line[len(line)-16:]
	if msg != "This is from log" {
		t.Errorf("log output \"%s\" does nto end with \"This is from log\"", msg)
	}
}
