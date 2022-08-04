package lrus

import (
	"github.com/sirupsen/logrus"
	"io"
)

type Lrus struct {
	Log *logrus.Logger
}

func New() *Lrus {
	return &Lrus{Log: logrus.New()}
}

func (l *Lrus) SetLevel(level string) {
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

func (l *Lrus) SetOutput(output io.Writer) {
	l.Log.SetOutput(output)
}

func (l *Lrus) Info(msg string) {
	l.Log.Infoln(msg)
}

func (l *Lrus) Error(msg string) {
	l.Log.Error(msg)
}

func (l *Lrus) Debug(msg string) {
	l.Log.Debug(msg)
}

func (l *Lrus) Warn(msg string) {
	l.Log.Warn(msg)
}
