package glog

import (
	"context"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var filePrefix = "file://"

// Logger is an interface of logging operations
type Logger interface {
	Infof(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Panicf(format string, v ...interface{})

	Infoc(ctx context.Context, format string, v ...interface{})
	Debugc(ctx context.Context, format string, v ...interface{})
	Warnc(ctx context.Context, format string, v ...interface{})
	Errorc(ctx context.Context, format string, v ...interface{})
	Panicc(ctx context.Context, format string, v ...interface{})

	WithField(field string, value interface{}) Logger
	Close() error
}

type closeFunc = func() error

// robusta is a logger implementation
type robusta struct {
	logger *logrus.Entry
	writer io.WriteCloser
}

// New return a new logger
// It will lookup configuration from environment variables for initialization
// LOG_FORMAT can be text/json
// LOG_OUTPUT can be a file by setting value to file:///path/to/logfile
// if LOG_OUTPUT is different to a file, it will be redirect to os.Stdout.
func New() *robusta {
	l := &robusta{}

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
func (l *robusta) Infof(format string, v ...interface{}) {
	l.logger.Infof(format, v...)
}

// Debugf print debug with format.
func (l *robusta) Debugf(format string, v ...interface{}) {
	l.logger.Debugf(format, v...)
}

// Warnf print warning with format.
func (l *robusta) Warnf(format string, v ...interface{}) {
	l.logger.Warnf(format, v...)
}

// Errorf print error with format.
func (l *robusta) Errorf(format string, v ...interface{}) {
	l.logger.Errorf(format, v...)
}

// Panicf panic with format.
func (l *robusta) Panicf(format string, v ...interface{}) {
	l.logger.Panicf(format, v...)
}

// Infoc print info log with context
func (l *robusta) Infoc(ctx context.Context, format string, v ...interface{}) {
	l.withContext(ctx).Infof(format, v...)
}

// Debugc print debug with context
func (l *robusta) Debugc(ctx context.Context, format string, v ...interface{}) {
	l.withContext(ctx).Debugf(format, v...)
}

// Warnc print warning with context
func (l *robusta) Warnc(ctx context.Context, format string, v ...interface{}) {
	l.withContext(ctx).Warnf(format, v...)
}

//Errorc print error with context
func (l *robusta) Errorc(ctx context.Context, format string, v ...interface{}) {
	l.withContext(ctx).Errorf(format, v...)
}

// Panicc panic with context
func (l *robusta) Panicc(ctx context.Context, format string, v ...interface{}) {
	l.withContext(ctx).Panicf(format, v...)
}

func (l *robusta) withContext(ctx context.Context) Logger {
	if requestID := ctx.Value("request_id"); requestID != nil {
		return l.WithField("request_id", requestID)
	}
	return l
}

// WithField return a new logger with field
func (l *robusta) WithField(field string, value interface{}) Logger {
	nl := l.logger.WithField(field, value)
	return &robusta{
		logger: nl,
	}
}

// Close close the underlying writer
func (l *robusta) Close() error {
	if l.writer != nil {
		return l.writer.Close()
	}
	return nil
}
