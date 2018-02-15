package utils

import "xframe/log"

type LogConfig struct {
	LogLevel string `json:"log_level"`
}

var (
	logConfig LogConfig
)

func InitLog(conf LogConfig) {
	logConfig = conf
	if logConfig.LogLevel == "" {
		logConfig.LogLevel = "DEBUG"
	}
	log.InitLogger("", "", "", 0, logConfig.LogLevel, "stdout")
}
