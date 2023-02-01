package api

import (
	"go.uber.org/zap"
	"os"
)

var logger *zap.SugaredLogger
var config APIConfiguration

func InitialiseAPI(apiConfig APIConfiguration, l *zap.SugaredLogger) {
	logger = l
	config = apiConfig
	err := os.MkdirAll(config.HarvestDataFolderPath, os.ModePerm)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
