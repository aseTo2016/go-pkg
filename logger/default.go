package logger

import "log"

type defaultLogger struct {
}

// InfoF ...
func (d defaultLogger) InfoF(template string, args ...interface{}) {
	log.Printf(template, args...)
}

// Info ...
func (d defaultLogger) Info(args ...interface{}) {
	log.Println(args...)
}

// DebugF ...
func (d defaultLogger) DebugF(template string, args ...interface{}) {
	log.Printf(template, args...)
}

// Debug ...
func (d defaultLogger) Debug(args ...interface{}) {
	log.Println(args...)
}

// ErrorF error must not be nil, if is nil ,please use Error
func (d defaultLogger) ErrorF(err error, template string, args ...interface{}) {
	log.Printf(template+", %s", append(args, err.Error())...)
}

// Error ...
func (d defaultLogger) Error(args ...interface{}) {
	log.Println(args...)
}

// WarnEf ...
func (d defaultLogger) WarnEf(err error, template string, args ...interface{}) {
	log.Printf(template+", %s", append(args, err.Error())...)
}

// WarnF ...
func (d defaultLogger) WarnF(template string, args ...interface{}) {
	log.Printf(template, args...)
}

// Warn ...
func (d defaultLogger) Warn(args ...interface{}) {
	log.Println(args...)
}
