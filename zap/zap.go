package zap

import (
	"github.com/aseTo2016/go-pkg/logger"
	"github.com/aseTo2016/go-pkg/pkg/yaml"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"runtime"
)

var logConfigPath string

var (
	lgr   *zap.Logger
	sugar *zap.SugaredLogger
)

// Init init config
func Init() error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}
	cfg.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	lgr, err = cfg.Build(zap.AddCallerSkip(2))
	if err != nil {
		return err
	}

	sugar = lgr.Sugar()
	logger.InitResource(new(adapter))

	logger.InfoF("log config path is %s", logConfigPath)

	return lgr.Sync()
}

func loadConfig() (*zap.Config, error) {
	logConfigPath = os.Getenv("log_config_path")
	if len(logConfigPath) == 0 {
		_, fileName, _, _ := runtime.Caller(1)
		logConfigPath = filepath.Join(filepath.Dir(fileName), "config", "log.yaml")
	}

	cfg := new(zap.Config)
	err := yaml.LoadYamlFromFile(logConfigPath, cfg)
	return cfg, err
}
