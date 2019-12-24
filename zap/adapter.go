package zap

type adapter struct {
}

// Infof ...
func (a *adapter) Infof(template string, args ...interface{}) {
	sugar.Infof(template, args...)
}

// Info ...
func (a *adapter) Info(args ...interface{}) {
	sugar.Info(args...)
}

// Debugf ...
func (a *adapter) Debugf(template string, args ...interface{}) {
	sugar.Debugf(template, args...)
}

// Debug ...
func (a *adapter) Debug(args ...interface{}) {
	sugar.Debug(args...)
}

// Errorf ...
func (a *adapter) Errorf(err error, template string, args ...interface{}) {
	sugar.Errorf(template+"%s", append(args, err.Error()))
}

// Error ...
func (a *adapter) Error(args ...interface{}) {
	sugar.Error(args...)
}
