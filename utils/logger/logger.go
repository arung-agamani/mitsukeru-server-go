package logger

import (
	"github.com/arung-agamani/mitsukeru-go/config"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func InitLogger() {
	logger = zap.Must(zap.NewProduction()).Sugar()

	if config.GetEnvironment() == "development" {
		logger = zap.Must(zap.NewDevelopment()).Sugar()
	}
	logger.WithOptions()
	zap.ReplaceGlobals(logger.Desugar())
}

func Info(args ...interface{}) {
	logger.Info(args)
}

func Infof(s string, args ...interface{}) {
	logger.Infof(s, args)
}

func Warn(args ...interface{})             { logger.Warn(args) }
func Warnf(s string, args ...interface{})  { logger.Warnf(s, args) }
func Error(args ...interface{})            { logger.Error(args) }
func Errorf(s string, args ...interface{}) { logger.Errorf(s, args) }
