package test

import (
	"fmt"
	"logl"
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
		t.Errorf("Wrong log level %s", level)
	}
	logl.Info("Test", "ing")

	logl.Debug("Aha!")
	logl.SetLevel(logl.WARN)
	logl.Debugf(">>%s<<", "dohh")

}
