package glog

import (
	"context"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/pthethanh/robusta/internal/pkg/kontext"
	"github.com/sirupsen/logrus"
)

var filePrefix = "file://"

// Logger is an interface of logging operations
type Logger interface {
	Info(format string)
	Debug(format string)
	Warn(format string)
	Error(format string)
	Panic(format string)

	Infof(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Panicf(format string, v ...interface{})

	Infoc(ctx context.Context, format string)
	Debugc(ctx context.Context, format string)
	Warnc(ctx context.Context, format string)
	Errorc(ctx context.Context, format string)
	Panicc(ctx context.Context, format string)

	Infocf(ctx context.Context, format string, v ...interface{})
	Debugcf(ctx context.Context, format string, v ...interface{})
	Warncf(ctx context.Context, format string, v ...interface{})
	Errorcf(ctx context.Context, format string, v ...interface{})
	Paniccf(ctx context.Context, format string, v ...interface{})

	WithField(field string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
	Close() error
}

type closeFunc = func() error

// GLog is a logger implementation
type GLog struct {
	logger *logrus.Entry
	writer io.WriteCloser
}

// New return a new logger
// It will lookup configuration from environment variables for initialization
// LOG_FORMAT can be text/json
// LOG_OUTPUT can be a file by setting value to file:///path/to/logfile
// if LOG_OUTPUT is different to a file, it will be redirect to os.Stdout.
func New() *GLog {
	l := &GLog{}

	logger := logrus.New()
	logger.SetFormatter(getFormater())
	logger.SetLevel(getLevel())

	out := getOutput()
	logger.SetOutput(out)

	l.writer = out
	l.logger = logrus.NewEntry(logger)
	return l
}

func getFormater() logrus.Formatter {
	var formatter logrus.Formatter
	formatter = &logrus.TextFormatter{
		TimestampFormat: time.RFC1123,
		FullTimestamp:   true,
	}
	if os.Getenv("LOG_FORMAT") == "json" {
		formatter = &logrus.JSONFormatter{
			TimestampFormat: time.RFC1123,
		}
	}
	return formatter
}

func getLevel() logrus.Level {
	lvl, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		lvl = logrus.DebugLevel
	}
	return lvl
}

func getOutput() io.WriteCloser {
	out := os.Getenv("LOG_OUTPUT")
	if strings.HasPrefix(out, filePrefix) {
		name := out[len(filePrefix):]
		f, err := os.Create(name)
		if err != nil {
			log.Printf("log: failed to create log file: %s, err: %v\n", name, err)
		}
		return f
	}
	return os.Stdout
}

// Infof print info with format.
func (l *GLog) Infof(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}

// Debugf print debug with format.
func (l *GLog) Debugf(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

// Warnf print warning with format.
func (l *GLog) Warnf(format string, v ...interface{}) {
	l.logger.Warnf(format, v...)
}

// Errorf print error with format.
func (l *GLog) Errorf(format string, v ...interface{}) {
	l.logger.Errorf(format, v...)
}

// Panicf panic with format.
func (l *GLog) Panicf(format string, v ...interface{}) {
	l.logger.Panicf(format, v...)
}

// Info print info with format.
func (l *GLog) Info(format string) {
	l.logger.Info(format)
}

// Debug print debug with format.
func (l *GLog) Debug(format string) {
	l.logger.Debug(format)
}

// Warn print warning with format.
func (l *GLog) Warn(format string) {
	l.logger.Warn(format)
}

// Error print error with format.
func (l *GLog) Error(format string) {
	l.logger.Error(format)
}

// Panic panic with format.
func (l *GLog) Panic(format string) {
	l.logger.Panicf(format)
}

// Infocf print info log with context
func (l *GLog) Infocf(ctx context.Context, format string, v ...interface{}) {
	l.withContext(ctx).Infof(format, v...)
}

// Debugcf print debug with context
func (l *GLog) Debugcf(ctx context.Context, format string, v ...interface{}) {
	l.withContext(ctx).Debugf(format, v...)
}

// Warncf print warning with context
func (l *GLog) Warncf(ctx context.Context, format string, v ...interface{}) {
	l.withContext(ctx).Warnf(format, v...)
}

//Errorcf print error with context
func (l *GLog) Errorcf(ctx context.Context, format string, v ...interface{}) {
	l.withContext(ctx).Errorf(format, v...)
}

// Paniccf panic with context
func (l *GLog) Paniccf(ctx context.Context, format string, v ...interface{}) {
	l.withContext(ctx).Panicf(format, v...)
}

// Infoc print info log with context
func (l *GLog) Infoc(ctx context.Context, format string) {
	l.withContext(ctx).Info(format)
}

// Debugc print debug with context
func (l *GLog) Debugc(ctx context.Context, format string) {
	l.withContext(ctx).Debug(format)
}

// Warnc print warning with context
func (l *GLog) Warnc(ctx context.Context, format string) {
	l.withContext(ctx).Warn(format)
}

//Errorc print error with context
func (l *GLog) Errorc(ctx context.Context, format string) {
	l.withContext(ctx).Error(format)
}

// Panicc panic with context
func (l *GLog) Panicc(ctx context.Context, format string) {
	l.withContext(ctx).Panicf(format)
}

func (l *GLog) withContext(ctx context.Context) Logger {
	if requestID := kontext.RequestIDFromContext(ctx); requestID != "" {
		return l.WithField("request_id", requestID)
	}
	return l
}

// WithField return a new logger with field
func (l *GLog) WithField(field string, value interface{}) Logger {
	nl := l.logger.WithField(field, value)
	return &GLog{
		logger: nl,
	}
}

// WithFields return a new logger with fields
func (l *GLog) WithFields(fields map[string]interface{}) Logger {
	nl := l.logger.WithFields(fields)
	return &GLog{
		logger: nl,
	}
}

// Close close the underlying writer
func (l *GLog) Close() error {
	if l.writer != nil {
		return l.writer.Close()
	}
	return nil
}
