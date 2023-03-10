package cmd

import (
	"github.com/antleaf/gopendoar/api"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
	"os"
)

// Configuration is the configuration for the application
type Configuration struct {
	APIConfig api.APIConfiguration `yaml:"opendoar"`
}

// Initialise initialises the application configuration
func (config *Configuration) Initialise(configFilePath string) error {
	configData, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(configData, config)
	if err != nil {
		return err
	}
	return err
}

// ConfigureZapSugarLogger configures the zap logger as a zap.SugaredLogger and returns a pointer to the logger and an error
func ConfigureZapSugarLogger(debugging bool) (*zap.SugaredLogger, error) {
	var zapLogger *zap.Logger
	var err error
	zapLogger, err = ConfigureZapLogger(debugging)
	return zapLogger.Sugar(), err
}

// ConfigureZapLogger configures the zap logger and returns a pointer to the logger and an error
func ConfigureZapLogger(debugging bool) (*zap.Logger, error) {
	level := zapcore.InfoLevel
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:  "message",
		LevelKey:    "level",
		TimeKey:     "",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
	}
	if debugging == true {
		level = zapcore.DebugLevel
		encoderConfig = zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			TimeKey:      "timestamp",
			EncodeTime:   zapcore.TimeEncoderOfLayout("Jan 02 15:04:05.000000000"),
			EncodeLevel:  zapcore.CapitalColorLevelEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
			//StacktraceKey: "stacktrace",
		}
	}
	zapConfig := zap.Config{
		Encoding:      "console",
		Level:         zap.NewAtomicLevelAt(level),
		OutputPaths:   []string{"stdout"},
		EncoderConfig: encoderConfig,
	}
	return zapConfig.Build()
}
