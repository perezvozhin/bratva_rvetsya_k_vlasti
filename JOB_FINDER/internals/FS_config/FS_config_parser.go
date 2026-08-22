package FS_config

type Config struct {
	PathFilesystem string `json:"PATH_TO_FILESYSTEM"`
	ApiKey         string `json:"API_KEY"`
	Port           string `json:"port"`
	Name           string `json:"name"`
}
