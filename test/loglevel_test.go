package test

import (
	"fmt"
	"loglevel"
	"testing"
)

func TestLogger(t *testing.T) {
	fmt.Println("Start")
	if level := loglevel.GetLevel(); level != loglevel.INFO {
		t.Errorf("Wrong loglevel %s", level)
	}
	loglevel.SetLevel(loglevel.DEBUG)
	if level := loglevel.GetLevel(); level != loglevel.DEBUG {
		t.Errorf("Wrong loglevel %s", level)
	}
	loglevel.Info("Test", "ing")

	loglevel.Debug("Aha!")
	loglevel.SetLevel(loglevel.WARN)
	loglevel.Debugf(">>%s<<", "dohh")

}
