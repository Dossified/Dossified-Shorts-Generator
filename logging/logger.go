// Provides logging functionality
package logging

import (
	"encoding/json"

	"go.uber.org/zap"

	"github.com/Dossified/Dossified-Shorts-Generator/config"
)

var zapLogger *zap.SugaredLogger

// Initializes the `zap` logger
func InitLogger() {
	logLevel := config.GetConfiguration().LogLevel
	rawJSON := []byte(`{
       "level": "` + logLevel + `",
       "encoding": "console",
       "outputPaths": ["stdout"],
       "errorOutputPaths": ["stderr"],
       "encoderConfig": {
         "messageKey": "message",
         "levelKey": "level",
         "levelEncoder": "lowercase"
       }
     }`)
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	zapLogger = logger.Sugar()
}

// Helper function to print `Info` log messages
func Info(message string, fields ...interface{}) {
	zapLogger.Infow(message, fields...)
}

// Helper function to print `Debug` log messages
func Debug(message string, fields ...interface{}) {
	zapLogger.Debugw(message, fields...)
}

// Helper function to print `Error` log messages
func Error(message string, fields ...interface{}) {
	zapLogger.Errorw(message, fields...)
}

// Helper function to print `Fatal` log messages
func Fatal(message string, fields ...interface{}) {
	zapLogger.Fatalw(message, fields...)
}
