package lrus

import (
	"github.com/sirupsen/logrus"
	"io"
)

type Logger struct {
	Log *logrus.Logger
}

func New() *Logger {
	return &Logger{Log: logrus.New()}
}

func (l *Logger) SetLevel(level string) {
	switch level {
	case "INFO":
		l.Log.Level = logrus.InfoLevel
	case "ERROR":
		l.Log.Level = logrus.ErrorLevel
	case "DEBUG":
		l.Log.Level = logrus.DebugLevel
	case "WARN":
		l.Log.Level = logrus.WarnLevel
	default:
		l.Log.Level = logrus.InfoLevel
	}
}

func (l *Logger) SetOutput(output io.Writer) {
	l.Log.SetOutput(output)
}

func (l *Logger) Info(msg string) {
	l.Log.Infoln(msg)
}

func (l *Logger) Error(msg string) {
	l.Log.Error(msg)
}

func (l *Logger) Debug(msg string) {
	l.Log.Debug(msg)
}

func (l *Logger) Warn(msg string) {
	l.Log.Warn(msg)
}
