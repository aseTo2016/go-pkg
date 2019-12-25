package zap

type adapter struct {
}

// InfoF ...
func (a *adapter) InfoF(template string, args ...interface{}) {
	sugar.Infof(template, args...)
}

// Info ...
func (a *adapter) Info(args ...interface{}) {
	sugar.Info(args...)
}

// DebugF ...
func (a *adapter) DebugF(template string, args ...interface{}) {
	sugar.Debugf(template, args...)
}

// Debug ...
func (a *adapter) Debug(args ...interface{}) {
	sugar.Debug(args...)
}

// ErrorF ...
func (a *adapter) ErrorF(err error, template string, args ...interface{}) {
	sugar.Errorf(template+"%s", append(args, err.Error()))
}

// Error ...
func (a *adapter) Error(args ...interface{}) {
	sugar.Error(args...)
}

// WarnEf ...
func (a *adapter) WarnEf(err error, template string, args ...interface{}) {
	sugar.Warnf(template+"%s", append(args, err.Error()))
}

// WarnF ...
func (a *adapter) WarnF(template string, args ...interface{}) {
	sugar.Warnf(template, args...)
}

// Warn ...
func (a *adapter) Warn(args ...interface{}) {
	sugar.Warn(args...)
}
