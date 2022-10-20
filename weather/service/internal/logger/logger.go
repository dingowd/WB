package logger

type Logger interface {
	Info(msg interface{})
	Error(msg interface{})
	Debug(msg interface{})
	Warn(msg interface{})
}
