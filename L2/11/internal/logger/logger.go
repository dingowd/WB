package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	Level string
	Log   *logrus.Logger
}

func New(l string) *Logger {
	return &Logger{Level: l, Log: logrus.New()}
}

func (l *Logger) SetLevel() {
	switch l.Level {
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
