package logger

// Logger interface
type Logger interface {
	Infof(string, ...interface{})
	Info(...interface{})
	Debugf(string, ...interface{})
	Debug(...interface{})
	Errorf(error, string, ...interface{})
	Error(...interface{})
}
