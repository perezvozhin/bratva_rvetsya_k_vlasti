package helper

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

/*
Функция переработана под открытие файла с резюме
Хелпер для сервиса пользователя
*/
func CheckDirectoryForCV(path string, logger *zap.SugaredLogger) *os.File {
	files, err := os.ReadDir(path)
	if err != nil {
		logger.Errorw("Error reading directory", "path", path, "error", err)
		return nil
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := filepath.Ext(file.Name())

		if ext == ".txt" || ext == ".pdf" {
			f, err := os.Open(filepath.Join(path, file.Name()))
			if err != nil {
				continue
			}
			logger.Debug("found file", "filename", file.Name())
			return f
		}
	}
	logger.Debug("No such file or directory")
	return nil
}
