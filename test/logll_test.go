package test

import (
	"bufio"
	"fmt"
	"logl"
	"math"
	"os"
	"testing"
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
func TestFileWrite(t *testing.T) {
	logl.SetLevel(logl.DEBUG)
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
	logl.Close()

	lf, err := os.Open("tmp/test.log")
	if err != nil {
		t.Fatal(err)
	}
	defer lf.Close()

	sc := bufio.NewScanner(lf)

	i := 0
	var l logl.Level

	for sc.Scan() {
		l = logl.Level(math.Mod(float64(i), 4.0) + 1)
		i++
		line := sc.Text()
		expected := fmt.Sprintf("[%s] Line %d", l, i)
		if line[33:] != expected {
			t.Errorf("line %d \"%s\" != \"%s\"", i, line[33:], expected)
		}
	}
}
