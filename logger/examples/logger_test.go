package examples

import (
	"os"
	"testing"
	"github.com/michaelzx/paladin-go/logger"
	"github.com/michaelzx/paladin-go/logger/stdlog"
)

type Foo struct {
	ID       int64
	SiteID   int64
	Title    string
	FormType int32
}

var foo = &Foo{
	ID:       1,
	SiteID:   2,
	Title:    "xxxxxx",
	FormType: 999,
}

func TestStdLog(t *testing.T) {
	logger.UseStdLog(os.Stderr, stdlog.LevelDebug, true)
	logger.Debug("Debug", *foo)
}

func TestZapLog(t *testing.T) {
	logger.UseZapLog(true, "")
	logger.Debug("Debug", foo)
}
