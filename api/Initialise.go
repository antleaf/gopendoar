package api

import "go.uber.org/zap"

var logger *zap.SugaredLogger
var config APIConfiguration

func InitialiseAPI(apiConfig APIConfiguration, l *zap.SugaredLogger) {
	logger = l
	config = apiConfig
}
