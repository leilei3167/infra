package zap

import (
	"github.com/leilei3167/infra/pkg/log"
	"go.uber.org/zap"
)

// Logger is a wrapper of  zap logger
type Logger struct {
	*zap.Logger
}

// config of zap logger
type opt struct {
}

func NewLogger(opts ...Option) *Logger {
	z := zap.NewExample()
	// 替换全局的logger为自定义的logger，简化使用
	_ = zap.ReplaceGlobals(z)
	return &Logger{z}
}

type Option func(*Logger)

func (l *Logger) Log(level log.Level, kvs ...interface{}) error {
	switch level {
	case log.LevelDebug:
		l.Sugar().Debug(kvs...)
	case log.LevelInfo:
		l.Sugar().Info(kvs...)
	case log.LevelWarn:
		l.Sugar().Warn(kvs...)
	case log.LevelError:
		l.Sugar().Error(kvs...)
	case log.LevelFatal:
		l.Sugar().Fatal(kvs...)
	}
	return nil
}
