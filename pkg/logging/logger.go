package logging

import (
	"os"
	"strings"

	"go.uber.org/zap"
)

func InitLogger() *zap.Logger {
	logger := zap.Must(zap.NewProduction())

	if strings.EqualFold(os.Getenv("APP_STATUS"), "debug") {
		logger = zap.Must(zap.NewDevelopment())
	}

	return logger
}
