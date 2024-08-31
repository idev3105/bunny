package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var logger = log.New()

// Init application logger
func init() {
	logger.Out = os.Stdout
	logger.Formatter = &log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
}

// ALL < TRACE < DEBUG < INFO < WARN < ERROR < FATAL < OFF
func SetLogLevel(level log.Level) {
	logger.SetLevel(level)
}

type Logger struct {
	entry         *log.Entry
	ComponentName string
	ServiceName   string
}

func New(componentName string, actionName string) *Logger {
	l := &Logger{ComponentName: componentName, ServiceName: actionName}
	l.entry = logger.WithFields(log.Fields{
		"ComponentName": componentName,
		"ActionName":    actionName,
	})
	return l
}

func (l *Logger) Debug(args ...any) {
	l.entry.Debug(args...)
}

func (l *Logger) Debugf(format string, args ...any) {
	l.entry.Debugf(format, args...)
}

func (l *Logger) Info(args ...any) {
	l.entry.Info(args...)
}

func (l *Logger) Infof(format string, args ...any) {
	l.entry.Infof(format, args...)
}

func (l *Logger) Warn(args ...any) {
	l.entry.Warn(args...)
}

func (l *Logger) Warnf(format string, args ...any) {
	l.entry.Warnf(format, args...)
}

func (l *Logger) Error(args ...any) {
	l.entry.Error(args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	l.entry.Errorf(format, args...)
}

func (l *Logger) Fatal(args ...any) {
	l.entry.Fatal(args...)
}

func (l *Logger) Fatalf(format string, args ...any) {
	l.entry.Fatalf(format, args...)
}
