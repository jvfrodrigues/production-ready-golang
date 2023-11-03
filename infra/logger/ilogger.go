package logger

type ILogger interface {
	Info(message string, args ...interface{})
	Error(message string, args ...interface{})
}
