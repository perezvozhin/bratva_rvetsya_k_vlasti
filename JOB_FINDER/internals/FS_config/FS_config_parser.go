package FS_config

type Config struct {
	PathStatic     string `json:"PATH_TO_STATIC"`
	PathFilesystem string `json:"PATH_TO_FILESYSTEM"`
	ApiKey         string `json:"API_KEY"`
	Port           string `json:"port"`
	Name           string `json:"name"`
}
