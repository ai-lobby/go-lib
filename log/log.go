package log

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"golang.org/x/exp/slog"
)

type Recorder interface {
	Info()
}

type defaultRecord struct {
	LogCaller bool
	logger    *slog.Logger
	msg       string
	now       slog.Attr
	caller    slog.Attr
}

func NewDefaultRecord() *defaultRecord {
	return &defaultRecord{
		logger: slog.New(slog.NewTextHandler(os.Stdout)),
	}
}

func (r *defaultRecord) SetLogger(h slog.Handler) error {
	r.logger = slog.New(h)
	return nil
}

func (r *defaultRecord) Info() {
	if r.LogCaller {
		_, f, l, _ := runtime.Caller(1)
		r.caller = slog.String("caller", fmt.Sprintf("%v:%v", f, l))
	}

	r.logger.Info(
		r.msg,
		r.now,
		r.caller,
	)
}

func (r *defaultRecord) Error() {
	if r.LogCaller {
		_, f, l, _ := runtime.Caller(1)
		r.caller = slog.String("caller", fmt.Sprintf("%v:%v", f, l))
	}

	r.logger.Error(
		r.msg,
		r.now,
		r.caller,
	)
}

func (r *defaultRecord) Debug() {
	if r.LogCaller {
		_, f, l, _ := runtime.Caller(1)
		r.caller = slog.String("caller", fmt.Sprintf("%v:%v", f, l))
	}

	r.logger.Debug(
		r.msg,
		r.now,
		r.caller,
	)
}

func (r *defaultRecord) Msg(msg string) *defaultRecord {
	r.msg = msg
	return r
}

func (r *defaultRecord) Now() *defaultRecord {
	r.now = slog.Time("now", time.Now())
	return r
}
