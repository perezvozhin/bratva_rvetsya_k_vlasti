package usr_service

import (
	"bufio"
	"os"

	"go.uber.org/zap"
)

/*
ПРОВЕРКА И ЧТЕНИЕ РЕЗЮМЕ ИЗ ФАЙЛОВОЙ СИСТЕМЫ
ОПЦИОНАЛЬНО ВЫДАЕМ И ТЕКСТ РЕЗЮМЕ И САМО РЕЗЮМЕ
(смотри DTO)!
*/
func (s *Usr_service) CheckReadCV() (RetvalsInCV, error) {
	var output RetvalsInCV
	file, err := os.Open(s.cv.Name())
	if err != nil {
		s.logger.Error("Error opening file", zap.String("filename", s.cv.Name()), zap.Error(err))
		return output, err
	}
	var CVText string
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s.logger.Debug(line)
		CVText += line
	}
	return output, err
}
