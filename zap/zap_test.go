package zap

import (
	"github.com/aseTo2016/go-pkg/logger"
	. "github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	err := Init()
	if Nil(t, err) {
		logger.Infof("---->, %s", "ok")
	}
}
