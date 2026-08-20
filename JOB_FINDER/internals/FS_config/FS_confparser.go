package FS_config

import (
	"encoding/json"
	"os"

	"go.uber.org/zap"
)

func Init(logs *zap.SugaredLogger) *Config {
	FS_fileText, err := os.ReadFile("../config/config.json")
	if err != nil {
		logs.Fatal("CANNOT READ CONFIG!")
		return nil
	}
	var config Config
	err = json.Unmarshal(FS_fileText, &config)
	if err != nil {
		logs.Fatal("CANNOT UNMARSHALL CONFIG!")
		return nil
	}
	logs.Infow("config", "config", config)
	return &config
}
