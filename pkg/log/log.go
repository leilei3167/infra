package log

type Level int8

const (
	LevelDebug Level = iota - 1
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

type Logger interface {
	Log(level Level, kvs ...interface{}) error
}
