package logger

// Infof ...
func Infof(template string, args ...interface{}) {
	instance.Infof(template, args...)
}

// Info ...
func Info(args ...interface{}) {
	instance.Info(args...)
}

// Debugf ...
func Debugf(template string, args ...interface{}) {
	instance.Debugf(template, args...)
}

// Debug ...
func Debug(args ...interface{}) {
	instance.Debug(args...)
}

// Errorf error must not be nil, if is nil ,please use Error
func Errorf(err error, template string, args ...interface{}) {
	instance.Errorf(err, template, args...)
}

// Error ...
func Error(args ...interface{}) {
	instance.Error(args...)
}
