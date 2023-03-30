package log

import (
	"os"
	"time"

	"golang.org/x/exp/slog"
)

type Recorder interface {
	Info()
}

type defaultRecord struct {
	logger    *slog.Logger
	msg       string
	timestamp time.Time
}

func NewDefaultRecord() *defaultRecord {
	return &defaultRecord{
		logger: slog.New(slog.NewTextHandler(os.Stdout)),
	}
}

func (r *defaultRecord) Info() {
	r.logger.Info(
		r.msg,
		slog.Time("timestamp", r.timestamp),
	)
}

func (r *defaultRecord) Msg(msg string) *defaultRecord {
	r.msg = msg
	return r
}

func (r *defaultRecord) Timestamp(timestamp time.Time) *defaultRecord {
	r.timestamp = timestamp
	return r
}
