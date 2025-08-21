package logger

import (
	"path/filepath"
	"runtime"
	"sync"

	"go.uber.org/zap"
)

var (
	log  *zap.Logger
	once sync.Once
)

func GetLogger() *zap.Logger {
	once.Do(func() {
		// Find directory of this source file (logger.go)
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			panic("Could not get current file path")
		}
		dir := filepath.Dir(filename)
		logFilePath := filepath.Join(dir, "app.log")

		config := zap.NewProductionConfig()

		config.OutputPaths = []string{logFilePath}
		config.ErrorOutputPaths = []string{logFilePath}
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		config.Encoding = "json"

		var err error
		log, err = config.Build(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
		if err != nil {
			panic(err)
		}
	})
	return log
}

func SugarLogger() *zap.SugaredLogger {
	return GetLogger().Sugar()
}
