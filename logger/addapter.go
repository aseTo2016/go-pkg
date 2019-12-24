package logger

import (
	"sync"
)

var (
	instance Logger = new(defaultLogger)
	once     sync.Once
)

// InitResource init log resource
func InitResource(l Logger) {
	once.Do(func() {
		instance = l
	})
}
