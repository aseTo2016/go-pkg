package logger

import "log"

type defaultLogger struct {
}

// Infof ...
func (d defaultLogger) Infof(template string, args ...interface{}) {
	log.Printf(template, args...)
}

// Info ...
func (d defaultLogger) Info(args ...interface{}) {
	log.Println(args...)
}

// Debugf ...
func (d defaultLogger) Debugf(template string, args ...interface{}) {
	log.Printf(template, args...)
}

// Debug ...
func (d defaultLogger) Debug(args ...interface{}) {
	log.Println(args...)
}

// Errorf error must not be nil, if is nil ,please use Error
func (d defaultLogger) Errorf(err error, template string, args ...interface{}) {
	log.Printf(template+", %s", append(args, err.Error())...)
}

// Error ...
func (d defaultLogger) Error(args ...interface{}) {
	log.Println(args...)
}
