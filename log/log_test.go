package log

import (
	"os"
	"testing"

	"golang.org/x/exp/slog"
)

func TestDefaultRecord_Log(t *testing.T) {
	r := NewDefaultRecord()
	r.LogCaller = true

	r.Msg("test msg").
		Now().
		Info()
}

func TestDefaultRecord_SetLogger(t *testing.T) {
	r := NewDefaultRecord()
	r.LogCaller = true
	r.SetLogger(slog.NewJSONHandler(os.Stdout))

	r.Msg("test msg").
		Now().
		Info()
}
