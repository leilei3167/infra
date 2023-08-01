package log

import (
	"io"
	std "log"
)

type StdLogger struct {
	*std.Logger
}

func NewStdLogger(w io.Writer) *StdLogger {
	return &StdLogger{
		Logger: std.New(w, "", 0),
	}
}

func (l *StdLogger) Log(level Level, kvs ...interface{}) error {
	l.Println(kvs...)
	return nil
}
