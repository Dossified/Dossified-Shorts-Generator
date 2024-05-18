package logging

import (
	"encoding/json"

	"go.uber.org/zap"

    "github.com/Dominique-Roth/Dossified-Shorts-Generator/config"
)

var zapLogger *zap.SugaredLogger

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

func Info(message string, fields ...interface{}) {
	zapLogger.Infow(message, fields...)
}

func Debug(message string, fields ...interface{}) {
	zapLogger.Debugw(message, fields...)
}

func Error(message string, fields ...interface{}) {
	zapLogger.Errorw(message, fields...)
}

func Fatal(message string, fields ...interface{}) {
	zapLogger.Fatalw(message, fields...)
}
