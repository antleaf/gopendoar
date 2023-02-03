// Package api provides datastructures and functions to interface with the OpenDOAR API
package api

import (
	"go.uber.org/zap"
	"os"
)

var logger *zap.SugaredLogger
var config APIConfiguration

// InitialiseAPI initialises the API package
func InitialiseAPI(apiConfig APIConfiguration, l *zap.SugaredLogger) {
	logger = l
	config = apiConfig
	err := os.MkdirAll(config.HarvestDataFolderPath, os.ModePerm)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
