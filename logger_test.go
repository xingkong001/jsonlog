package jsonlog

import (
	"github.com/funny/utest"
	"testing"
	"time"
)

func Test_SwitchByDay(t *testing.T) {
	log, err := New(Config{
		Dir:      ".",
		Switcher: DAY_SWITCHER,
		FileType: ".log",
	})
	utest.NotError(t, err)
	log.Log(M{"Time": time.Now()})
	log.Log(M{"Time": time.Now()})
	log.Log(M{"Time": time.Now()})
	log.Close()
}

func Test_SwitchByHours(t *testing.T) {
	log, err := New(Config{
		Dir:      ".",
		Switcher: HOURS_SWITCHER,
		FileType: ".log",
	})
	utest.NotError(t, err)
	log.Log(M{"Time": time.Now()})
	log.Log(M{"Time": time.Now()})
	log.Log(M{"Time": time.Now()})
	log.Close()
}
