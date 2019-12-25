package logger

// Logger interface
type Logger interface {
	InfoF(string, ...interface{})
	Info(...interface{})
	DebugF(string, ...interface{})
	Debug(...interface{})
	ErrorF(error, string, ...interface{})
	Error(...interface{})
	WarnEf(error, string, ...interface{})
	WarnF(string, ...interface{})
	Warn(...interface{})
}
