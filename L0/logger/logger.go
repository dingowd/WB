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

func (l *Logger) SetLevel() logrus.Level {
	switch l.Level {
	case "INFO":
		return logrus.InfoLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "DEBUG":
		return logrus.DebugLevel
	case "WARN":
		return logrus.WarnLevel
	default:
		return logrus.InfoLevel
	}
}

func (l *Logger) SetOutput(output io.Writer) {
	l.Log.Out = output
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
