package FS_config

import "go.uber.org/zap"

type Config struct {
	TimeStamp      string `json:"TIME_STAMP"`
	PathFilesystem string `json:"PATH_TO_FILESYSTEM"`
	ApiKey         string `json:"API_KEY"`
	Port           string `json:"port"`
	Name           string `json:"name"`
}

func Init(logger *zap.SugaredLogger) *Config {
	var Cfg *Config
	return Cfg
}
