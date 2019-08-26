package log

import (
	"context"
)

type (
	Logger interface {
		// Infof print info with format.
		Infof(format string, v ...interface{})

		// Debugf print debug with format.
		Debugf(format string, v ...interface{})

		// Warnf print warning with format.
		Warnf(format string, v ...interface{})

		// Errorf print error with format.
		Errorf(format string, v ...interface{})

		// Panicf panic with format.
		Panicf(format string, v ...interface{})

		// Info print info.
		Info(v ...interface{})

		// Debug print debug.
		Debug(v ...interface{})

		// Warn print warning.
		Warn(v ...interface{})

		// Error print error.
		Error(v ...interface{})

		// Panic panic.
		Panic(v ...interface{})

		WithField(key string, val interface{}) Logger

		WithFields(fields Fields) Logger
	}
	// Fields is alias of map
	Fields = map[string]interface{}

	// context key
	contextKey string
)

const (
	loggerKey  contextKey = contextKey("logger_key")
	filePrefix            = "file://"
)

var (
	root Logger
)

// Root return default logger instance
func Root() Logger {
	if root == nil {
		root = newGlog()
	}
	return root
}

// Init init the root logger with fields
func Init(fields Fields) {
	root = newGlogWithFields(fields)
}

// NewWithPrefix return new logger with context
func NewWithPrefix(key, val string) Logger {
	return newGlogWithField(key, val)
}

// New return new logger with context
func New(ctx Fields) Logger {
	return newGlogWithFields(ctx)
}

// NewContext return a new logger context
func NewContext(ctx context.Context, logger Logger) context.Context {
	if logger == nil {
		logger = Root()
	}
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext get logger form context
func FromContext(ctx context.Context) Logger {
	if ctx == nil {
		return Root()
	}
	if logger, ok := ctx.Value(loggerKey).(Logger); ok {
		return logger
	}
	return Root()
}

// Infof print info with format.
func Infof(format string, v ...interface{}) {
	Root().Infof(format, v...)
}

// Debugf print debug with format.
func Debugf(format string, v ...interface{}) {
	Root().Debugf(format, v...)
}

// Warnf print warning with format.
func Warnf(format string, v ...interface{}) {
	Root().Warnf(format, v...)
}

// Errorf print error with format.
func Errorf(format string, v ...interface{}) {
	Root().Errorf(format, v...)
}

// Panicf panic with format.
func Panicf(format string, v ...interface{}) {
	Root().Panicf(format, v...)
}

// WithFields return a new logger entry with fields
func WithFields(fields Fields) Logger {
	return Root().WithFields(fields)
}

// WithContext return a logger from the given context
func WithContext(ctx context.Context) Logger {
	return FromContext(ctx)
}
