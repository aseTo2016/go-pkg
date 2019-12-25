package logger

// InfoF ...
func InfoF(template string, args ...interface{}) {
	instance.InfoF(template, args...)
}

// Info ...
func Info(args ...interface{}) {
	instance.Info(args...)
}

// DebugF ...
func DebugF(template string, args ...interface{}) {
	instance.DebugF(template, args...)
}

// Debug ...
func Debug(args ...interface{}) {
	instance.Debug(args...)
}

// ErrorF error must not be nil, if is nil ,please use Error
func ErrorF(err error, template string, args ...interface{}) {
	instance.ErrorF(err, template, args...)
}

// Error ...
func Error(args ...interface{}) {
	instance.Error(args...)
}

// WarnEf ...
func WarnEf(err error, template string, args ...interface{}) {
	instance.WarnEf(err, template, args)
}

// WarnEf ...
func WarnF(template string, args ...interface{}) {
	instance.WarnF(template, args...)
}

// Warn ...
func Warn(args ...interface{}) {
	instance.Warn(args...)
}
